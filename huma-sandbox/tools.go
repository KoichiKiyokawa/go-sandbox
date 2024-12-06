//go:build tools

package tools

import (
	_ "github.com/air-verse/air"
	_ "github.com/go-jet/jet/v2/cmd/jet"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt"
)
