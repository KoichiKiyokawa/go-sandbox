run-with-log:
	LOG_QUERIES=y PHOTON_GO_LOG=info go run cmd/api/main.go

generate:
	@go run github.com/steebchen/prisma-client-go generate

db-push:
	@go run github.com/steebchen/prisma-client-go db push

seed:
	@go run script/seed.go

reset-and-seed:
	@rm -rf dev.db
	@make db-push
	@make seed