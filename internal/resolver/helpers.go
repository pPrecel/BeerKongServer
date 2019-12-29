package resolver

import "github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"

func intToInt32(value *int) *int32 {
	var converted *int32
	if value != nil {
		tmp := int32(*value)
		converted = &tmp
	}
	return converted
}
func (r *Resolver) getAdmins() []string {
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

func (r *Resolver) chooseMatchWhereUniqueInput(match *prisma.Match) * prisma.MatchWhereUniqueInput {
	if match == nil {
		return nil
	}

	if match.ID != "" {
		return &prisma.MatchWhereUniqueInput{ID: &match.ID}
	}

	return nil
}
func (r *Resolver) chooseUserWhereUniqueInput(user *prisma.User) * prisma.UserWhereUniqueInput {
	if user == nil {
		return nil
	}

	if user.ID != "" {
		return &prisma.UserWhereUniqueInput{ID: &user.ID}
	}
	if user.Name != "" {
		return &prisma.UserWhereUniqueInput{Name: &user.Name}
	}
	if user.Sub != "" {
		return &prisma.UserWhereUniqueInput{Sub: &user.Sub}
	}

	return nil
}
func (r *Resolver) chooseTeamWhereUniqueInput(team *prisma.Team) * prisma.TeamWhereUniqueInput {
	if team == nil {
		return nil
	}

	if team.ID != "" {
		return &prisma.TeamWhereUniqueInput{ID: &team.ID}
	}
	if team.Name != "" {
		return &prisma.TeamWhereUniqueInput{Name: &team.Name}
	}

	return nil
}
func (r *Resolver) chooseLeagueWhereUniqueInput(league *prisma.League) * prisma.LeagueWhereUniqueInput {
	if league == nil {
		return nil
	}

	if league.ID != "" {
		return &prisma.LeagueWhereUniqueInput{ID: &league.ID}
	}
	if league.Name != "" {
		return &prisma.LeagueWhereUniqueInput{Name: &league.Name}
	}

	return nil
}

func (r *Resolver) fillUserWhereUniqueInput(user *prisma.UserWhereUniqueInput) *prisma.UserWhereUniqueInput {
	if user == nil {
		return nil
	}
	if user.ID != nil {
		return &prisma.UserWhereUniqueInput{ID: user.ID}
	}
	if user.Sub != nil {
		return &prisma.UserWhereUniqueInput{Sub: user.Sub}
	}
	if user.Name != nil {
		return &prisma.UserWhereUniqueInput{Name: user.Name}
	}
	return nil
}

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
func buildMatchesParams(where *prisma.MatchWhereInput, orderBy *prisma.MatchOrderByInput, skip *int, after *string, before *string, first *int, last *int) *prisma.MatchesParams {
	return &prisma.MatchesParams{
		Where:   where,
		OrderBy: orderBy,
		Skip:    intToInt32(skip),
		After:   after,
		Before:  before,
		First:   intToInt32(first),
		Last:    intToInt32(last),
	}
}

