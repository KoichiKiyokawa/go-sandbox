package integration

import (
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestIntegration_CreateAccount(t *testing.T) {
	_, api := humatest.New(t)

	i := do.New()
	do.ProvideValue(i, db)
	do.ProvideValue(i, api)
	do.Provide(i, infra.NewDao)
	do.Provide(i, uiapi.NewHandler)
	handler := do.MustInvoke[*uiapi.Handler](i)
	handler.RegisterAll()

	resp := api.Post("/account")
	assert.JSONEq(t, resp.Body, `
	foo: 1
}
	`)
}
