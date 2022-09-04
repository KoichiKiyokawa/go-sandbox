package resolver

import (
	mock_service "fx-di/service/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func withDIContainer(t *testing.T, testFunc any) {
	fxtest.New(
		t,
		fx.Provide(
			func() *gomock.Controller { return gomock.NewController(t) },
			mock_service.NewMockUserService,
		),
		fx.Invoke(testFunc),
	).RequireStart()
}
