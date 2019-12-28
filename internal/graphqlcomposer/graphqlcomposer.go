package graphqlcomposer

import (
	"github.com/pPrecel/BeerKongServer/pkg/graphql/resolver"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

type Composer interface {
	Resolver(user *prisma.User) *resolver.Resolver
}

type composer struct {
	client *prisma.Client
}

func New(client *prisma.Client) Composer {
	return &composer{client: client}
}

func (c *composer) Resolver(user *prisma.User) *resolver.Resolver {
	return resolver.New(c.client, user)
}
