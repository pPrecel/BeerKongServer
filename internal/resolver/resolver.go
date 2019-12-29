package resolver

import (
	"context"
	"fmt"
	"github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
	prisma "github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	prismaClient *prisma.Client
	user         *prisma.User
}

func New(prismaClient *prisma.Client, user *prisma.User) *Resolver {
	return &Resolver{prismaClient: prismaClient, user: user}
}

func (r *Resolver) League() generated.LeagueResolver {
	return &leagueResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Team() generated.TeamResolver {
	return &teamResolver{r}
}
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Match() generated.MatchResolver {
	return &matchResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) updateUserPoints(ctx context.Context, user *prisma.User, pointsToAdd int) (*prisma.User, error) {
	finalPoints := user.Points + *intToInt32(&pointsToAdd)
	return r.prismaClient.UpdateUser(prisma.UserUpdateParams{
		Data: prisma.UserUpdateInput{
			Points: &finalPoints,
		},
		Where: *r.chooseUserWhereUniqueInput(user),
	}).Exec(ctx)
}
func (r *mutationResolver) updateTeamPoints(ctx context.Context, team *prisma.Team, pointsToAdd int) (*prisma.Team, error) {
	finalPoints := team.Points + *intToInt32(&pointsToAdd)
	return r.prismaClient.UpdateTeam(prisma.TeamUpdateParams{
		Data: prisma.TeamUpdateInput{
			Points: &finalPoints,
		},
		Where: *r.chooseTeamWhereUniqueInput(team),
	}).Exec(ctx)
}
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
	}).Exec(ctx)
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
		return nil, fmt.Errorf("User %s is a member of this or another team in this league", data.Name)
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
	one := 1
	winner, _ := r.prismaClient.Match(where).Winner().Exec(ctx)
	if winner != nil {
		return nil, errors.New(fmt.Sprintf("This event has been updated"))
	}
	owner, err := r.prismaClient.Match(where).League().Owner().Exec(ctx)
	if err != nil {
		return nil, err
	}
	if !r.checkAccessWithUser(owner.Sub) {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}

	league, err := r.prismaClient.Match(where).League().Exec(ctx)
	if err != nil {
		return nil, err
	}

	user1, err := r.prismaClient.Match(where).User1().Exec(ctx)
	if err != nil {
		return nil, err
	}
	user1, err = r.updateUserPoints(ctx, user1, data.User1points)
	if err != nil {
		return nil, err
	}

	team1, err := r.prismaClient.Teams(&prisma.TeamsParams{
		Where: &prisma.TeamWhereInput{
			And: []prisma.TeamWhereInput{
				prisma.TeamWhereInput{
					League: &prisma.LeagueWhereInput{Name: &league.Name},
				},
				prisma.TeamWhereInput{
					UsersSome: &prisma.UserWhereInput{ID: &user1.ID},
				},
			},
		},
		First: intToInt32(&one),
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.updateTeamPoints(ctx, &team1[0], data.User1points)
	if err != nil {
		return nil, err
	}

	user2, err := r.prismaClient.Match(where).User2().Exec(ctx)
	if err != nil {
		return nil, err
	}
	r.updateUserPoints(ctx, user2, data.User2points)

	team2, err := r.prismaClient.Teams(&prisma.TeamsParams{
		Where: &prisma.TeamWhereInput{
			And: []prisma.TeamWhereInput{
				prisma.TeamWhereInput{
					League: &prisma.LeagueWhereInput{Name: &league.Name},
				},
				prisma.TeamWhereInput{
					UsersSome: &prisma.UserWhereInput{ID: &user2.ID},
				},
			},
		},
		First: intToInt32(&one),
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.updateTeamPoints(ctx, &team2[0], data.User2points)
	if err != nil {
		return nil, err
	}

	var winnerPoints *int32
	if data.User1points > data.User2points {
		winner = user1
		winnerPoints = intToInt32(&data.User1points)
	} else {
		winner = user2
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

type queryResolver struct{ *Resolver }

func (r *queryResolver) League(ctx context.Context, where prisma.LeagueWhereUniqueInput) (*prisma.League, error) {
	return r.prismaClient.League(where).Exec(ctx)
}
func (r *queryResolver) Leagues(ctx context.Context, where *prisma.LeagueWhereInput, orderBy *prisma.LeagueOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.League, error) {
	data, err := r.prismaClient.Leagues(buildLeaguesParams(where, orderBy, skip, after, before, first, last)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	dataPointers := make([]*prisma.League, len(data))
	for index, _ := range data {
		dataPointers[index] = &data[index]
	}

	return dataPointers, nil
}
func (r *queryResolver) Team(ctx context.Context, where prisma.TeamWhereUniqueInput) (*prisma.Team, error) {
	return r.prismaClient.Team(where).Exec(ctx)
}
func (r *queryResolver) Teams(ctx context.Context, where *prisma.TeamWhereInput, orderBy *prisma.TeamOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.Team, error) {
	data, err := r.prismaClient.Teams(buildTeamsParams(where, orderBy, skip, after, before, first, last)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	dataPointers := make([]*prisma.Team, len(data))
	for index, _ := range data {
		dataPointers[index] = &data[index]
	}

	return dataPointers, nil
}
func (r *queryResolver) User(ctx context.Context, where prisma.UserWhereUniqueInput) (*prisma.User, error) {
	return r.prismaClient.User(where).Exec(ctx)
}
func (r *queryResolver) Users(ctx context.Context, where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.User, error) {
	data, err := r.prismaClient.Users(buildUsersParams(where, orderBy, skip, after, before, first, last)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	dataPointers := make([]*prisma.User, len(data))
	for index, _ := range data {
		dataPointers[index] = &data[index]
	}

	return dataPointers, nil
}
func (r *queryResolver) Match(ctx context.Context, where prisma.MatchWhereUniqueInput) (*prisma.Match, error) {
	return r.prismaClient.Match(where).Exec(ctx)
}
func (r *queryResolver) Matches(ctx context.Context, where *prisma.MatchWhereInput, orderBy *prisma.MatchOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.Match, error) {
	data, err := r.prismaClient.Matches(buildMatchesParams(where, orderBy, skip, after, before, first, last)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	dataPointers := make([]*prisma.Match, len(data))
	for index, _ := range data {
		dataPointers[index] = &data[index]
	}

	return dataPointers, nil
}

type leagueResolver struct{ *Resolver }

func (r *leagueResolver) Teams(ctx context.Context, obj *prisma.League) ([]prisma.Team, error) {
	return r.prismaClient.League(prisma.LeagueWhereUniqueInput{
		ID: &obj.ID,
	}).Teams(nil).Exec(ctx)
}
func (r *leagueResolver) Users(ctx context.Context, obj *prisma.League) ([]prisma.User, error) {
	return r.prismaClient.League(prisma.LeagueWhereUniqueInput{
		ID: &obj.ID,
	}).Users(nil).Exec(ctx)
}
func (r *leagueResolver) Owner(ctx context.Context, obj *prisma.League) (*prisma.User, error) {
	return r.prismaClient.League(prisma.LeagueWhereUniqueInput{
		ID: &obj.ID,
	}).Owner().Exec(ctx)
}

type teamResolver struct{ *Resolver }

func (r *teamResolver) League(ctx context.Context, obj *prisma.Team) (*prisma.League, error) {
	return r.prismaClient.Team(prisma.TeamWhereUniqueInput{
		ID: &obj.ID,
	}).League().Exec(ctx)
}
func (r *teamResolver) Users(ctx context.Context, obj *prisma.Team) ([]prisma.User, error) {
	return r.prismaClient.Team(prisma.TeamWhereUniqueInput{
		ID: &obj.ID,
	}).Users(nil).Exec(ctx)
}
func (r *teamResolver) Owner(ctx context.Context, obj *prisma.Team) (*prisma.User, error) {
	return r.prismaClient.Team(prisma.TeamWhereUniqueInput{
		ID: &obj.ID,
	}).Owner().Exec(ctx)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Matches(ctx context.Context, obj *prisma.User) ([]prisma.Match, error) {
	return r.prismaClient.User(*r.chooseUserWhereUniqueInput(obj)).Matches(nil).Exec(ctx)
}
func (r *userResolver) Teams(ctx context.Context, obj *prisma.User) ([]prisma.Team, error) {
	return r.prismaClient.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Teams(nil).Exec(ctx)
}
func (r *userResolver) Leagues(ctx context.Context, obj *prisma.User) ([]prisma.League, error) {
	return r.prismaClient.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Leagues(nil).Exec(ctx)
}
func (r *userResolver) OwnedTeams(ctx context.Context, obj *prisma.User) ([]prisma.Team, error) {
	return r.prismaClient.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).OwnedTeams(nil).Exec(ctx)
}
func (r *userResolver) OwnedLeagues(ctx context.Context, obj *prisma.User) ([]prisma.League, error) {
	return r.prismaClient.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).OwnedLeagues(nil).Exec(ctx)
}

type matchResolver struct{ *Resolver }

func (r *matchResolver) League(ctx context.Context, obj *prisma.Match) (*prisma.League, error) {
	return r.prismaClient.Match(*r.chooseMatchWhereUniqueInput(obj)).League().Exec(ctx)
}
func (r *matchResolver) User1(ctx context.Context, obj *prisma.Match) (*prisma.User, error) {
	return r.prismaClient.Match(*r.chooseMatchWhereUniqueInput(obj)).User1().Exec(ctx)
}
func (r *matchResolver) User2(ctx context.Context, obj *prisma.Match) (*prisma.User, error) {
	return r.prismaClient.Match(*r.chooseMatchWhereUniqueInput(obj)).User2().Exec(ctx)
}
func (r *matchResolver) Winner(ctx context.Context, obj *prisma.Match) (*prisma.User, error) {
	return r.prismaClient.Match(*r.chooseMatchWhereUniqueInput(obj)).Winner().Exec(ctx)
}
