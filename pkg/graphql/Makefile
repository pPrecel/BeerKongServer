.EXPORT_ALL_VARIABLES:
.PHONY: generate-graphql delete-graphql

delete-graphql:
	@rm -rf server generate.go gqlgen.yml models_gen.go resolver.go

generate-graphql: delete-graphql
	@go run github.com/99designs/gqlgen init
	@rm -rf server
	@echo "done"
