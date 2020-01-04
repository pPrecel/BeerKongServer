PRISMA_SECRET ?= PrismaSecret

.EXPORT_ALL_VARIABLES:
.PHONY: prisma-deploy prisma-generate prisma-local

prisma-deploy:
	@echo "Prisma deploy:"
	@prisma deploy

prisma-generate:
	@echo "Prisma generate manifest:"
	@prisma generate

prisma-local:
	@echo "Prisma locally run:"
	@docker-compose up -d
