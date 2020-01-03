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
func (r *resolver) getAdmins() []string {
	return []string{
		"111923199050409291371",
		"103308660506651951126",
	}
}
func (r *resolver) checkAccessWithUser(sub string) bool {
	for _, adminSub := range r.getAdmins() {
		if sub == adminSub {
			return true
		}
	}

	if r.user == nil {
		return false
	}
	if r.isSameUser(r.user, &prisma.User{Sub: sub}) {
		return false
	}

	return true
}

func (r *resolver) chooseMatchWhereUniqueInput(match *prisma.Match) *prisma.MatchWhereUniqueInput {
	if match == nil {
		return nil
	}

	if match.ID != "" {
		return &prisma.MatchWhereUniqueInput{ID: &match.ID}
	}

	return nil
}
func (r *resolver) chooseTeamWhereUniqueInput(team *prisma.Team) *prisma.TeamWhereUniqueInput {
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
func (r *resolver) chooseLeagueWhereUniqueInput(league *prisma.League) *prisma.LeagueWhereUniqueInput {
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
func (r *resolver) chooseUserWhereUniqueInput(user *prisma.User) *prisma.UserWhereUniqueInput {
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

func (r *resolver) isSameMatch(match1 *prisma.Match, match2 *prisma.Match) bool {
	if match1 == nil || match2 == nil {
		return false
	}
	if match1.ID == match2.ID {
		return true
	}
	return false
}
func (r *resolver) isSameLeague(league1 *prisma.Team, league2 *prisma.Team) bool {
	if league1 == nil || league2 == nil {
		return false
	}
	if league1.ID == league2.ID {
		return true
	}
	return false
}
func (r *resolver) isSameTeam(team1 *prisma.Team, team2 *prisma.Team) bool {
	if team1 == nil || team2 == nil {
		return false
	}
	if team1.ID == team2.ID {
		return true
	}
	return false
}
func (r *resolver) isSameUser(user1 *prisma.User, user2 *prisma.User) bool {
	if user1 == nil || user2 == nil {
		return false
	}
	if user1.ID == user2.ID {
		return true
	}
	if user1.Sub == user2.Sub {
		return true
	}
	if user1.Name == user2.Name {
		return true
	}
	return false
}

func (r *resolver) fillUserWhereUniqueInput(user *prisma.UserWhereUniqueInput) *prisma.UserWhereUniqueInput {
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
