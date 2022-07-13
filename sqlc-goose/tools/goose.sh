#!/bin/bash
set -euxo pipefail
cd "$(dirname "$0")"

go mod tidy
GOOSE_DRIVER=postgres go run github.com/pressly/goose/v3/cmd/goose --dir ../sql/migrations "$@"
