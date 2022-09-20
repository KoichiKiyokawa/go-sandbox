DB_URL="host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
TEST_DB_URL="host=localhost port=5432 user=postgres dbname=test password=postgres sslmode=disable"
TEST_FILES=`go list ./app/... ./test/... | grep -v mock`

.PHONY: help dev build create-schema generate db-exec migrate migrate-reset seed seed-reset lint test

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Run the app in development mode
	docker compose down
	docker compose up -d
	DB_URL=$(DB_URL) DB_DEBUG=$(DB_DEBUG) go run github.com/cosmtrek/air --build.cmd="make build"

build: ## build the application
	go build -o tmp/main cmd/api/main.go

create-schema: ## create schema of ent. e.g. NAME=User make create-schema
	@echo "Creating schema..."
	go run -mod=mod entgo.io/ent/cmd/ent init $(NAME)

generate: ## generate schema, resolver, etc...
	go generate ./...

db-exec: ## exec command for db. e.g. make db-exec COMMAND="select * from users"
	docker compose exec postgres psql -U postgres -c "$(COMMAND)"

migrate: ## db migrate
	DB_URL=$(DB_URL) go run script/migrate.go

migrate-reset: ## db migrate with reset
	make db-exec DBNAME=postgres COMMAND="DROP SCHEMA public CASCADE;"
	make db-exec DBNAME=postgres COMMAND="CREATE SCHEMA public;"
	make migrate

seed: ## db seed
	DB_URL=$(DB_URL) go run script/seed.go

seed-reset: migrate-reset seed ## db seed with reset

lint: ## run lint
	go vet ./...

test: ## run tests
	make db-exec COMMAND="DROP DATABASE IF EXISTS test;"
	make db-exec COMMAND="CREATE DATABASE test;"
	DB_URL=$(TEST_DB_URL) go test -v -cover -coverprofile=tmp/cover.out -json $(TEST_FILES) 2>&1 | go run github.com/haveyoudebuggedit/gotestfmt/v2/cmd/gotestfmt
	go tool cover -html=tmp/cover.out -o tmp/cover.html
