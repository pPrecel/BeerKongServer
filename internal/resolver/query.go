package resolver

import (
	"context"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

type queryResolver struct{ *resolver }

func (r *queryResolver) League(ctx context.Context, where prisma.LeagueWhereUniqueInput) (*prisma.League, error) {
	return r.prismaClient.League(where).Exec(ctx)
}
func (r *queryResolver) LeaguePoints(ctx context.Context, where prisma.LeagueWhereUniqueInput) (*int, error) {
	tr := true
	points := 0
	matches, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			League:     &prisma.LeagueWhereInput{ID: where.ID},
			IsFinished: &tr,
			IsRanked:   &tr,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		points += int(match.User1points)
		points += int(match.User2points)
	}
	return &points, nil
}
func (r *queryResolver) LeagueFinishedMatches(ctx context.Context, where prisma.LeagueWhereUniqueInput) (*int, error) {
	tr := true
	matches, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			League:     &prisma.LeagueWhereInput{ID: where.ID},
			IsRanked:   &tr,
			IsFinished: &tr,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	matchesLen := len(matches)
	return &matchesLen, nil
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
func (r *queryResolver) TeamPoints(ctx context.Context, where prisma.TeamWhereUniqueInput) (*int, error) {
	tr := true
	points := 0
	matches1, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			IsRanked:   &tr,
			IsFinished: &tr,
			User1: &prisma.UserWhereInput{TeamsSome: &prisma.TeamWhereInput{
				ID: where.ID,
			}},
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	matches2, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			IsRanked:   &tr,
			IsFinished: &tr,
			User1: &prisma.UserWhereInput{TeamsSome: &prisma.TeamWhereInput{
				ID: where.ID,
			}},
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	for _, match := range matches1 {
		points += int(match.User1points)
	}
	for _, match := range matches2 {
		points += int(match.User2points)
	}

	return &points, nil
}
func (r *queryResolver) TeamFinishedMatches(ctx context.Context, where prisma.TeamWhereUniqueInput) (*int, error) {
	tr := true
	matches, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			IsRanked:   &tr,
			IsFinished: &tr,
			Or: []prisma.MatchWhereInput{
				prisma.MatchWhereInput{User1: &prisma.UserWhereInput{TeamsSome: &prisma.TeamWhereInput{
					ID: where.ID,
				}}},
				prisma.MatchWhereInput{User2: &prisma.UserWhereInput{TeamsSome: &prisma.TeamWhereInput{
					ID: where.ID,
				}}},
			},
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	matchesLen := len(matches)
	return &matchesLen, nil
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
func (r *queryResolver) UserPoints(ctx context.Context, where prisma.UserWhereUniqueInput) (*int, error) {
	tr := true
	points := 0
	user1Matches, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			User1:      &prisma.UserWhereInput{ID: where.ID},
			IsFinished: &tr,
			IsRanked:   &tr,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	user2Matches, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			User2:      &prisma.UserWhereInput{ID: where.ID},
			IsFinished: &tr,
			IsRanked:   &tr,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	for _, match := range user1Matches {
		points += int(match.User1points)
	}
	for _, match := range user2Matches {
		points += int(match.User2points)
	}

	return &points, nil
}
func (r *queryResolver) UserPointsInLeague(ctx context.Context, where prisma.UserWhereUniqueInput, in prisma.LeagueWhereUniqueInput) (*int, error) {
	tr:=true
	points:=0
	matches1, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where:   &prisma.MatchWhereInput{
			League: &prisma.LeagueWhereInput{ID: in.ID},
			User1: &prisma.UserWhereInput{ID: where.ID},
			IsFinished: &tr,
			IsRanked: &tr,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	matches2, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where:   &prisma.MatchWhereInput{
			League: &prisma.LeagueWhereInput{ID: in.ID},
			User2: &prisma.UserWhereInput{ID: where.ID},
			IsFinished: &tr,
			IsRanked: &tr,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	for _, match := range matches1 {
		points += int(match.User1points)
	}
	for _, match := range matches2 {
		points += int(match.User2points)
	}

	return &points, nil
}
func (r *queryResolver) UserFinishedMatches(ctx context.Context, where prisma.UserWhereUniqueInput) (*int, error) {
	tr := true
	matches, err := r.prismaClient.Matches(&prisma.MatchesParams{
		Where: &prisma.MatchWhereInput{
			IsRanked:   &tr,
			IsFinished: &tr,
			Or: []prisma.MatchWhereInput{
				prisma.MatchWhereInput{User1: &prisma.UserWhereInput{ID: where.ID}},
				prisma.MatchWhereInput{User2: &prisma.UserWhereInput{ID: where.ID}},
			},
		},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	matchesLen := len(matches)
	return &matchesLen, nil
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

type leagueResolver struct{ *resolver }

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

type teamResolver struct{ *resolver }

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

type userResolver struct{ *resolver }

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

type matchResolver struct{ *resolver }

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
	unique := *r.chooseMatchWhereUniqueInput(obj)
	data, _ := r.prismaClient.Match(unique).Winner().Exec(ctx)
	return data, nil
}
