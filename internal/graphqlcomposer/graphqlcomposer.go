package graphqlcomposer

import (
	graphql "github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

type Composer interface {
	Resolver(user *prisma.User) *graphql.Resolver
}

type composer struct {
	client *prisma.Client
}

func New(client *prisma.Client) Composer {
	return &composer{client: client}
}

func (c *composer) Resolver(user *prisma.User) *graphql.Resolver {
	return graphql.New(c.client, user)
}
