.EXPORT_ALL_VARIABLES:
.PHONY: generate-graphql

generate-graphql:
	make -C pkg/graphql/ generate-graphql