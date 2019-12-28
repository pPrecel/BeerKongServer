// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	prisma "github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

type LeagueCreateInput struct {
	Description string                        `json:"description"`
	Name        string                        `json:"name"`
	Users       []prisma.UserWhereUniqueInput `json:"users"`
}

type MatchCreateInput struct {
	Expiration string                         `json:"expiration"`
	IsRanked   bool                           `json:"isRanked"`
	League     *prisma.LeagueWhereUniqueInput `json:"league"`
	User1      *prisma.UserWhereUniqueInput   `json:"user1"`
	User2      *prisma.UserWhereUniqueInput   `json:"user2"`
}

type TeamCreateInput struct {
	ID          *string                        `json:"id"`
	Description string                         `json:"description"`
	Name        string                         `json:"name"`
	League      *prisma.LeagueWhereUniqueInput `json:"league"`
}

type UserCreateInput struct {
	ID      *string `json:"id"`
	Name    string  `json:"name"`
	Sub     string  `json:"sub"`
	Picture string  `json:"picture"`
}
