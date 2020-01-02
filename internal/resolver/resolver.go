package resolver

import (
	"github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
	prisma "github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver interface {
	League() generated.LeagueResolver
	Query() generated.QueryResolver
	Team() generated.TeamResolver
	User() generated.UserResolver
	Mutation() generated.MutationResolver
	Match() generated.MatchResolver
}

type resolver struct {
	prismaClient *prisma.Client
	user         *prisma.User
}

func New(prismaClient *prisma.Client, user *prisma.User) Resolver {
	return &resolver{prismaClient: prismaClient, user: user}
}

func (r *resolver) League() generated.LeagueResolver {
	return &leagueResolver{r}
}
func (r *resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *resolver) Team() generated.TeamResolver {
	return &teamResolver{r}
}
func (r *resolver) User() generated.UserResolver {
	return &userResolver{r}
}
func (r *resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *resolver) Match() generated.MatchResolver {
	return &matchResolver{r}
}
