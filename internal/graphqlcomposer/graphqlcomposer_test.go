package graphqlcomposer_test

import (
	"github.com/onsi/gomega"
	"github.com/pPrecel/BeerKongServer/internal/graphqlcomposer"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
	"testing"
)

func TestComposer_Resolver(t *testing.T) {
	t.Run("when everything is ok", func(t *testing.T){
		//given
		g := gomega.NewWithT(t)
		prismaClient := prisma.Client{}

		//when
		composer := graphqlcomposer.New(&prismaClient)
		resolver := composer.Resolver(nil)

		//then
		g.Expect(resolver).NotTo(gomega.BeNil())
	})
}