//go:build tools
// +build tools

package tools

import (
	_ "ariga.io/atlas/cmd/atlas"
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
	_ "github.com/matryer/moq"
)
