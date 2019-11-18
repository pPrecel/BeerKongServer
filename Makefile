PORT ?= 80

.EXPORT_ALL_VARIABLES:
.PHONY: generate-graphql update-graphql go-run

generate-graphql:
	@make -C pkg/graphql/ generate-graphql

update-graphql:
	@make -C pkg/graphql/ update-graphql

go-run:
	@export PORT={PORT}
	@go run cmd/main.go

