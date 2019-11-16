.EXPORT_ALL_VARIABLES:
.PHONY: generate-graphql update-graphql

generate-graphql:
	@make -C pkg/graphql/ generate-graphql

update-graphql:
	@make -C pkg/graphql/ update-graphql
