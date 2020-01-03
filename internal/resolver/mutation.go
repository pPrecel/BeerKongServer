package resolver

import (
	"context"
	"fmt"
	"github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type mutationResolver struct{ *resolver }

func (r *mutationResolver) createRankMatch(ctx context.Context, data generated.MatchCreateInput) (*prisma.Match, error) {
	falseValue := false
	leagueOwner, err := r.prismaClient.League(*data.League).Owner().Exec(ctx)
	if err != nil {
		return nil, err
	}
	if !r.checkAccessWithUser(leagueOwner.Sub) {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}
	match, err := r.prismaClient.CreateMatch(prisma.MatchCreateInput{
		PlannedAt:  data.PlannedAt,
		IsRanked:   &data.IsRanked,
		IsFinished: &falseValue,
		League:     prisma.LeagueCreateOneWithoutMatchesInput{Connect: data.League},
		User1:      prisma.UserCreateOneInput{Connect: r.fillUserWhereUniqueInput(data.User1)},
		User2:      prisma.UserCreateOneInput{Connect: r.fillUserWhereUniqueInput(data.User2)},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.prismaClient.UpdateUser(prisma.UserUpdateParams{
		Data: prisma.UserUpdateInput{
			Matches: &prisma.MatchUpdateManyInput{
				Connect: []prisma.MatchWhereUniqueInput{*r.chooseMatchWhereUniqueInput(match)},
			},
		},
		Where: *r.fillUserWhereUniqueInput(data.User1),
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.prismaClient.UpdateUser(prisma.UserUpdateParams{
		Data: prisma.UserUpdateInput{
			Matches: &prisma.MatchUpdateManyInput{
				Connect: []prisma.MatchWhereUniqueInput{*r.chooseMatchWhereUniqueInput(match)},
			},
		},
		Where: *r.fillUserWhereUniqueInput(data.User2),
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return match, err
}
func (r *mutationResolver) isUserInLeague(ctx context.Context, where prisma.TeamWhereUniqueInput, data prisma.UserWhereUniqueInput) (*prisma.League, bool) {
	league, err := r.prismaClient.Team(where).League().Exec(ctx)
	if err != nil {
		return nil, true
	}

	users, err := r.prismaClient.League(*r.chooseLeagueWhereUniqueInput(league)).Users(nil).Exec(ctx)
	if err != nil {
		return league, true
	}

	for _, user := range users {
		if user.ID == *data.ID {
			return league, true
		}
	}
	return league, false
}

func (r *mutationResolver) CreateLeague(ctx context.Context, data generated.LeagueCreateInput) (*prisma.League, error) {
	if r.user == nil {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}
	return r.prismaClient.CreateLeague(prisma.LeagueCreateInput{
		Description: data.Description,
		Name:        data.Name,
		Users: &prisma.UserCreateManyWithoutLeaguesInput{
			Connect: data.Users,
		},
		Owner: prisma.UserCreateOneWithoutOwnedLeaguesInput{
			Connect: &prisma.UserWhereUniqueInput{
				Sub: &r.user.Sub,
			},
		},
	}).Exec(ctx)
}
func (r *mutationResolver) DeleteLeague(ctx context.Context, where prisma.LeagueWhereUniqueInput) (*prisma.League, error) {
	owner, err := r.prismaClient.League(where).Owner().Exec(ctx)
	if err != nil {
		return nil, err
	}
	if !r.checkAccessWithUser(owner.Sub) {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}

	return r.prismaClient.DeleteLeague(where).Exec(ctx)
}
func (r *mutationResolver) CreateTeam(ctx context.Context, data generated.TeamCreateInput) (*prisma.Team, error) {
	if r.user == nil {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}
	return r.prismaClient.CreateTeam(prisma.TeamCreateInput{
		Description: data.Description,
		Name:        data.Name,
		League: prisma.LeagueCreateOneWithoutTeamsInput{
			Connect: data.League,
		},
		Owner: prisma.UserCreateOneWithoutOwnedTeamsInput{
			Connect: &prisma.UserWhereUniqueInput{
				Sub: &r.user.Sub,
			},
		},
		Users: &prisma.UserCreateManyWithoutTeamsInput{
			Connect: []prisma.UserWhereUniqueInput{
				*r.chooseUserWhereUniqueInput(r.user),
			},
		},
	}).Exec(ctx)
}
func (r *mutationResolver) RemoveUserFromTeam(ctx context.Context, data generated.TeamCreateInput) (*prisma.Team, error) {
	return nil, nil
}
func (r *mutationResolver) DeleteTeam(ctx context.Context, where prisma.TeamWhereUniqueInput) (*prisma.Team, error) {
	owner, err := r.prismaClient.Team(where).Owner().Exec(ctx)
	if err != nil {
		return nil, err
	}
	if !r.checkAccessWithUser(owner.Sub) {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}

	return r.prismaClient.DeleteTeam(where).Exec(ctx)
}
func (r *mutationResolver) LoginOrRegisterUser(ctx context.Context, data generated.UserCreateInput) (*prisma.User, error) {
	if user, err := r.prismaClient.User(prisma.UserWhereUniqueInput{
		Sub: &data.Sub,
	}).Exec(ctx); err == nil {
		updated, err := r.prismaClient.UpdateUser(prisma.UserUpdateParams{
			Data: prisma.UserUpdateInput{
				Name:    &data.Name,
				Sub:     &data.Sub,
				Picture: &data.Picture,
			},
			Where: prisma.UserWhereUniqueInput{
				Sub: &data.Sub,
			},
		}).Exec(ctx)
		if err != nil {
			logrus.Infof("User not updated: %s", err.Error())
			return user, err
		}
		return updated, nil
	} else {
		return r.prismaClient.CreateUser(prisma.UserCreateInput{
			Name:    data.Name,
			Sub:     data.Sub,
			Picture: data.Picture,
		}).Exec(ctx)
	}
}
func (r *mutationResolver) AddUserToTeam(ctx context.Context, where prisma.TeamWhereUniqueInput, data prisma.UserWhereUniqueInput) (*prisma.Team, error) {
	league, isIn := r.isUserInLeague(ctx, where, data)
	if isIn {
		return nil, fmt.Errorf("User %s is a member of this or another team in this league", *data.ID)
	}

	team, err := r.prismaClient.UpdateTeam(prisma.TeamUpdateParams{
		Data: prisma.TeamUpdateInput{Users: &prisma.UserUpdateManyWithoutTeamsInput{
			Connect: []prisma.UserWhereUniqueInput{
				data,
			},
		}},
		Where: where,
	}).Exec(ctx)
	if err != nil {
		return team, err
	}

	_, err = r.prismaClient.UpdateUser(prisma.UserUpdateParams{
		Data: prisma.UserUpdateInput{
			Leagues: &prisma.LeagueUpdateManyWithoutUsersInput{
				Connect: []prisma.LeagueWhereUniqueInput{
					prisma.LeagueWhereUniqueInput{ID: &league.ID},
				},
			},
		},
		Where: data,
	}).Exec(ctx)
	if err != nil {
		return team, err
	}

	return team, nil
}
func (r *mutationResolver) DeleteUser(ctx context.Context) (*prisma.User, error) {
	if r.user == nil {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}

	return r.prismaClient.DeleteUser(prisma.UserWhereUniqueInput{
		Sub: &r.user.Sub,
	}).Exec(ctx)
}
func (r *mutationResolver) CreateMatch(ctx context.Context, data generated.MatchCreateInput) (*prisma.Match, error) {
	if data.IsRanked {
		return r.createRankMatch(ctx, data)
	}

	return nil, fmt.Errorf("Unexpected error xd")
}
func (r *mutationResolver) EndMatch(ctx context.Context, where prisma.MatchWhereUniqueInput, data generated.MatchEndInput) (*prisma.Match, error) {
	tr := true
	match, _ := r.prismaClient.Match(where).Exec(ctx)
	if match.IsFinished != false {
		return nil, errors.New(fmt.Sprintf("This event has been updated"))
	}
	owner, err := r.prismaClient.Match(where).League().Owner().Exec(ctx)
	if err != nil {
		return nil, err
	}
	if !r.checkAccessWithUser(owner.Sub) {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}

	var winnerPoints *int32
	var winner *prisma.User
	if data.User1points > data.User2points {
		winner, err = r.prismaClient.Match(where).User1().Exec(ctx)
		if err != nil {
			return nil, err
		}
		winnerPoints = intToInt32(&data.User1points)
	} else {
		winner, err = r.prismaClient.Match(where).User2().Exec(ctx)
		if err != nil {
			return nil, err
		}
		winnerPoints = intToInt32(&data.User2points)
	}

	return r.prismaClient.UpdateMatch(prisma.MatchUpdateParams{
		Data: prisma.MatchUpdateInput{
			User1points: intToInt32(&data.User1points),
			User2points: intToInt32(&data.User2points),
			Winner: &prisma.UserUpdateOneInput{
				Connect: r.chooseUserWhereUniqueInput(winner),
			},
			WinnerPoints: winnerPoints,
			IsFinished:   &tr,
		},
		Where: where,
	}).Exec(ctx)
}
func (r *mutationResolver) DeleteMatch(ctx context.Context, where prisma.MatchWhereUniqueInput) (*prisma.Match, error) {
	league, err := r.prismaClient.Match(where).League().Exec(ctx)
	if err != nil {
		return nil, err
	}
	owner, err := r.prismaClient.League(*r.chooseLeagueWhereUniqueInput(league)).Owner().Exec(ctx)
	if err != nil {
		return nil, err
	}
	if !r.checkAccessWithUser(owner.Sub) {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}

	return r.prismaClient.DeleteMatch(where).Exec(ctx)
}
