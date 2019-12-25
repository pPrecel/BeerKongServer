package generated

import (
	"context"
	prisma "github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	prismaClient *prisma.Client
}

func New(prismaClient *prisma.Client) *Resolver {
	return &Resolver{prismaClient: prismaClient}
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

func intToInt32(value *int) *int32 {
	var converted *int32
	if value != nil {tmp := int32(*value); converted = &tmp}
	return converted
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
	for index, value := range data {
		dataPointers[index] = &value
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
	for index, value := range data {
		dataPointers[index] = &value
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
	for index, value := range data {
		dataPointers[index] = &value
	}

	return dataPointers, nil
}

type leagueResolver struct{ *Resolver }
func buildTeamsParamsExec(where *prisma.TeamWhereInput, orderBy *prisma.TeamOrderByInput, skip *int, after *string, before *string, first *int, last *int) *prisma.TeamsParamsExec {
	return &prisma.TeamsParamsExec{
		Where:   where,
		OrderBy: orderBy,
		Skip:    intToInt32(skip),
		After:   after,
		Before:  before,
		First:   intToInt32(first),
		Last:    intToInt32(last),
	}
}

func (r *leagueResolver) Teams(ctx context.Context, obj *prisma.League, where *prisma.TeamWhereInput, orderBy *prisma.TeamOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]prisma.Team, error) {
	return r.prismaClient.League(prisma.LeagueWhereUniqueInput{
		ID:   &obj.ID,
	}).Teams(buildTeamsParamsExec(where, orderBy, skip, after, before, first, last)).Exec(ctx)
}
func (r *leagueResolver) Users(ctx context.Context, obj *prisma.League, where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]prisma.User, error) {
	panic("not implemented")
}
func (r *leagueResolver) Owner(ctx context.Context, obj *prisma.League) (*prisma.User, error) {
	data, err := r.prismaClient.League(prisma.LeagueWhereUniqueInput{
		ID:   &obj.ID,
	}).Owner().Exec(ctx)

	return data, err
}

type teamResolver struct{ *Resolver }
func (r *teamResolver) League(ctx context.Context, obj *prisma.Team) (*prisma.League, error) {
	panic("not implemented")
}
func (r *teamResolver) Users(ctx context.Context, obj *prisma.Team, where *prisma.UserWhereInput, orderBy *prisma.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]prisma.User, error) {
	panic("not implemented")
}
func (r *teamResolver) Owner(ctx context.Context, obj *prisma.Team) (*prisma.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }
func (r *userResolver) Teams(ctx context.Context, obj *prisma.User, where *prisma.TeamWhereInput, orderBy *prisma.TeamOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]prisma.Team, error) {
	panic("not implemented")
}
func (r *userResolver) Leagues(ctx context.Context, obj *prisma.User, where *prisma.LeagueWhereInput, orderBy *prisma.LeagueOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]prisma.League, error) {
	panic("not implemented")
}
func (r *userResolver) OwnedTeams(ctx context.Context, obj *prisma.User, where *prisma.TeamWhereInput, orderBy *prisma.TeamOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]prisma.Team, error) {
	panic("not implemented")
}
func (r *userResolver) OwnedLeagues(ctx context.Context, obj *prisma.User, where *prisma.LeagueWhereInput, orderBy *prisma.LeagueOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]prisma.League, error) {
	panic("not implemented")
}
