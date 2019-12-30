package resolver

import (
	"github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
	prisma "github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
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
