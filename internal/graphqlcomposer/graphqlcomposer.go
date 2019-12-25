package graphqlcomposer

import (
	"github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
)

type Composer interface {
	Resolver() *generated.Resolver
}

type composer struct {
	client *prisma.Client
}

func New(client *prisma.Client) Composer {
	return &composer{client: client}
}

func (c *composer) Resolver() *generated.Resolver {
	return generated.New(c.client)
}
