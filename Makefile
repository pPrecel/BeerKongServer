PORT ?= 80
PRISMA_SECRET ?= PrismaSecret

.EXPORT_ALL_VARIABLES:
.PHONY: go-fmt go-run go-test go-get gqlgen-regenerate prisma-deploy prisma-generate prisma-local

go-fmt:
	@go fmt ./...

go-get:
	@echo "Go get:"
	@go get -u ./...

go-test: go-get
	@echo "Run tests:"
	@go test ./internal/...

go-run: go-test
	@echo "Run server:"
	@export PORT={PORT}
	@go run cmd/main.go

gqlgen-regenerate:
	@make -C pkg/graphql/ gqlgen-regenerate

prisma-deploy:
	@make -C pkg/prisma/ prisma-deploy
	@make prisma-generate

prisma-generate:
	@make -C pkg/prisma/ prisma-generate

prisma-local:
	@make -C pkg/prisma/ prisma-local
