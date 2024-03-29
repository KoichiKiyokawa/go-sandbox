// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"fx-di/app/service"
	"fx-di/ent"
	"sync"
)

// Ensure, that UserServiceMock does implement service.UserService.
// If this is not the case, regenerate this file with moq.
var _ service.UserService = &UserServiceMock{}

// UserServiceMock is a mock implementation of service.UserService.
//
//	func TestSomethingThatUsesUserService(t *testing.T) {
//
//		// make and configure a mocked service.UserService
//		mockedUserService := &UserServiceMock{
//			FindAllFunc: func(ctx context.Context) ([]*ent.User, error) {
//				panic("mock out the FindAll method")
//			},
//			FindOneFunc: func(ctx context.Context, id int) (*ent.User, error) {
//				panic("mock out the FindOne method")
//			},
//		}
//
//		// use mockedUserService in code that requires service.UserService
//		// and then make assertions.
//
//	}
type UserServiceMock struct {
	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx context.Context) ([]*ent.User, error)

	// FindOneFunc mocks the FindOne method.
	FindOneFunc func(ctx context.Context, id int) (*ent.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FindOne holds details about calls to the FindOne method.
		FindOne []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int
		}
	}
	lockFindAll sync.RWMutex
	lockFindOne sync.RWMutex
}

// FindAll calls FindAllFunc.
func (mock *UserServiceMock) FindAll(ctx context.Context) ([]*ent.User, error) {
	if mock.FindAllFunc == nil {
		panic("UserServiceMock.FindAllFunc: method is nil but UserService.FindAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	mock.lockFindAll.Unlock()
	return mock.FindAllFunc(ctx)
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//
//	len(mockedUserService.FindAllCalls())
func (mock *UserServiceMock) FindAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockFindAll.RLock()
	calls = mock.calls.FindAll
	mock.lockFindAll.RUnlock()
	return calls
}

// FindOne calls FindOneFunc.
func (mock *UserServiceMock) FindOne(ctx context.Context, id int) (*ent.User, error) {
	if mock.FindOneFunc == nil {
		panic("UserServiceMock.FindOneFunc: method is nil but UserService.FindOne was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindOne.Lock()
	mock.calls.FindOne = append(mock.calls.FindOne, callInfo)
	mock.lockFindOne.Unlock()
	return mock.FindOneFunc(ctx, id)
}

// FindOneCalls gets all the calls that were made to FindOne.
// Check the length with:
//
//	len(mockedUserService.FindOneCalls())
func (mock *UserServiceMock) FindOneCalls() []struct {
	Ctx context.Context
	ID  int
} {
	var calls []struct {
		Ctx context.Context
		ID  int
	}
	mock.lockFindOne.RLock()
	calls = mock.calls.FindOne
	mock.lockFindOne.RUnlock()
	return calls
}
