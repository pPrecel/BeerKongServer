PORT ?= 80

.EXPORT_ALL_VARIABLES:
.PHONY: generate-graphql update-graphql go-run deploy-prisma generate-prisma

generate-graphql:
	@make -C pkg/graphql/ generate-new-graphql

update-graphql:
	@make -C pkg/graphql/ update-graphql

go-run:
	@export PORT={PORT}
	@go run cmd/main.go

deploy-prisma:
	@make -C pkg/prisma/ deploy-prisma

generate-prisma:
	@make -C pkg/prisma/ generate-prisma
