.EXPORT_ALL_VARIABLES:
.PHONY: generate-graphql delete-graphql update-graphql

delete-graphql:
	@rm -rf server generate.go gqlgen.yml models_gen.go resolver.go

generate-new-graphql: delete-graphql
	@go run github.com/99designs/gqlgen init -v
	@rm -rf server
	@echo "done"

update-graphql:
	@rm -rf resolver.go
	@go run github.com/99designs/gqlgen -v
