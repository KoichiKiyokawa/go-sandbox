DEV_DB_URL="host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
TEST_DB_URL="host=localhost port=5432 user=postgres dbname=test password=postgres sslmode=disable"

.PHONY: build
build: ## build the application
	go build -o tmp/main cmd/api/main.go

.PHONY: dev
dev: ## run dev server
	DB_URL=$(DEV_DB_URL) DB_DEBUG=$(DB_DEBUG) go run github.com/cosmtrek/air --build.cmd="make build"

.PHONY: create-schema
create-schema: ## create schema of ent e.g. NAME=User make create-schema
	@echo "Creating schema..."
	go run -mod=mod entgo.io/ent/cmd/ent init $(NAME)

.PHONY: generate
generate: ## generate schema, resolver, etc...
	go generate ./...

.PHONY: migrate
migrate: ## db migrate
	DB_URL=$(DEV_DB_URL) go run script/migrate.go

.PHONY: migrate-reset
migrate-reset: ## db migrate with reset
	RESET=1 make migrate

.PHONY: seed
seed: ## db seed
	DB_URL=$(DEV_DB_URL) go run script/seed.go

.PHONY: seed-reset
seed-reset: ## db seed with reset
	@make migrate-reset
	@make seed

.PHONY: lint
lint: ## run lint
	go vet ./...

.PHONY: test
test: ## run tests
	docker compose exec postgres psql -U postgres -c "DROP DATABASE IF EXISTS test;"
	docker compose exec postgres psql -U postgres -c "CREATE DATABASE test;"
	DB_URL=$(TEST_DB_URL) go test -v -json ./... 2>&1 | go run github.com/haveyoudebuggedit/gotestfmt/v2/cmd/gotestfmt

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'