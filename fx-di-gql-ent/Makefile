.PHONY: create-schema
create-schema: ## create schema of ent e.g. NAME=User make create-schema
	@echo "Creating schema..."
	go run -mod=mod entgo.io/ent/cmd/ent init $(NAME)

.PHONY: generate
generate: ## generate schema, resolver, etc...
	go generate ./...

.PHONY: migrate
migrate: ## db migrate
	go run script/migrate.go

.PHONY: migrate-reset
migrate-reset: ## db migrate with reset
	rm dev.db && make migrate

.PHONY: seed ## db seed
seed:
	go run script/seed.go

.PHONY: seed ## db seed with reset
seed-reset:
	make migrate-reset && go run script/seed.go

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'