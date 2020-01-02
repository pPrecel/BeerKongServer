package graphqlcomposer

import (
	resolver2 "github.com/pPrecel/BeerKongServer/internal/resolver"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

type Composer interface {
	Resolver(user *prisma.User) resolver2.Resolver
}

type composer struct {
	client *prisma.Client
}

func New(client *prisma.Client) Composer {
	return &composer{client: client}
}

func (c *composer) Resolver(user *prisma.User) resolver2.Resolver {
	return resolver2.New(c.client, user)
}
