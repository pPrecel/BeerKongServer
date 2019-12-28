package generated

import (
	"context"
	"fmt"
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

func (r *Resolver) League() LeagueResolver {
	return &leagueResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Team() TeamResolver {
	return &teamResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Match() MatchResolver{
	return &matchResolver{r}
}

func intToInt32(value *int) *int32 {
	var converted *int32
	if value != nil {
		tmp := int32(*value)
		converted = &tmp
	}
	return converted
}

type mutationResolver struct{ *Resolver }

func (r *Resolver) getAdmins() []string{
	return []string{
		"111923199050409291371",
		"103308660506651951126",
	}
}
func (r *Resolver) checkAccessWithUser(sub string) bool {
	for _, adminSub := range r.getAdmins() {
		if sub == adminSub {
			return true
		}
	}

	if r.user == nil {
		return false
	}
	if r.user.Sub != sub {
		return false
	}

	return true
}

func (r *mutationResolver) CreateLeague(ctx context.Context, data LeagueCreateInput) (*prisma.League, error) {
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
func (r *mutationResolver) CreateTeam(ctx context.Context, data TeamCreateInput) (*prisma.Team, error) {
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
func (r *mutationResolver) CreateUser(ctx context.Context, data UserCreateInput) (*prisma.User, error) {
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
func (r *mutationResolver) DeleteUser(ctx context.Context) (*prisma.User, error) {
	if r.user == nil {
		return nil, errors.New(fmt.Sprintf("You don't have permission to finish this operation"))
	}

	return r.prismaClient.DeleteUser(prisma.UserWhereUniqueInput{
		Sub: &r.user.Sub,
	}).Exec(ctx)
}
func (r *mutationResolver) CreateMatch(ctx context.Context, data MatchCreateInput) (*prisma.Match, error) {
	return nil, nil
}
func (r *mutationResolver) DeleteMatch(ctx context.Context, where prisma.MatchWhereUniqueInput) (*prisma.Match, error) {
	return nil, nil
}

type queryResolver struct{ *Resolver }

func buildLeaguesParams(where *prisma.LeagueWhereInput, orderBy *prisma.LeagueOrderByInput, skip *int, after *string, before *string, first *int, last *int) *prisma.LeaguesParams {
	return &prisma.LeaguesParams{
		Where:   where,
		OrderBy: orderBy,
		Skip:    intToInt32(skip),
		After:   after,
		Before:  before,
		First:   intToInt32(first),
		Last:    intToInt32(last),
	}
}
func buildTeamsParams(where *prisma.TeamWhereInput, orderBy *prisma.TeamOrderByInput, skip *int, after *string, before *string, first *int, last *int) *prisma.TeamsParams {
	return &prisma.TeamsParams{
		Where:   where,
		OrderBy: orderBy,
		Skip:    intToInt32(skip),
		After:   after,
		Before:  before,
		First:   intToInt32(first),
		Last:    intToInt32(last),
	}
}
func buildUsersParams(where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) *prisma.UsersParams {
	return &prisma.UsersParams{
		Where:   where,
		OrderBy: orderBy,
		Skip:    intToInt32(skip),
		After:   after,
		Before:  before,
		First:   intToInt32(first),
		Last:    intToInt32(last),
	}
}

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
func (r *queryResolver) Match(ctx context.Context, where prisma.MatchWhereUniqueInput) (*prisma.Match, error){
	return nil, nil
}
func (r *queryResolver) Matches(ctx context.Context, where *prisma.MatchWhereInput, orderBy *prisma.MatchOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]*prisma.Match, error){
	return nil, nil
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

func (r *matchResolver) League(ctx context.Context, obj *prisma.Match) (*prisma.League, error){
	return nil, nil
}
func (r *matchResolver) User1(ctx context.Context, obj *prisma.Match) (*prisma.User, error){
	return nil, nil
}
func (r *matchResolver) User2(ctx context.Context, obj *prisma.Match) (*prisma.User, error){
	return nil, nil
}
func (r *matchResolver) Winner(ctx context.Context, obj *prisma.Match) (*prisma.User, error){
	return nil, nil
}
