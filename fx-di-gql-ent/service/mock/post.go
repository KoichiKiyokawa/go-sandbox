// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"fx-di/ent"
	"fx-di/service"
	"sync"
)

// Ensure, that PostServiceMock does implement service.PostService.
// If this is not the case, regenerate this file with moq.
var _ service.PostService = &PostServiceMock{}

// PostServiceMock is a mock implementation of service.PostService.
//
//	func TestSomethingThatUsesPostService(t *testing.T) {
//
//		// make and configure a mocked service.PostService
//		mockedPostService := &PostServiceMock{
//			FindAllFunc: func(ctx context.Context) ([]*ent.Post, error) {
//				panic("mock out the FindAll method")
//			},
//			FindOneFunc: func(ctx context.Context, id int) (*ent.Post, error) {
//				panic("mock out the FindOne method")
//			},
//		}
//
//		// use mockedPostService in code that requires service.PostService
//		// and then make assertions.
//
//	}
type PostServiceMock struct {
	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx context.Context) ([]*ent.Post, error)

	// FindOneFunc mocks the FindOne method.
	FindOneFunc func(ctx context.Context, id int) (*ent.Post, error)

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
func (mock *PostServiceMock) FindAll(ctx context.Context) ([]*ent.Post, error) {
	if mock.FindAllFunc == nil {
		panic("PostServiceMock.FindAllFunc: method is nil but PostService.FindAll was just called")
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
//	len(mockedPostService.FindAllCalls())
func (mock *PostServiceMock) FindAllCalls() []struct {
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
func (mock *PostServiceMock) FindOne(ctx context.Context, id int) (*ent.Post, error) {
	if mock.FindOneFunc == nil {
		panic("PostServiceMock.FindOneFunc: method is nil but PostService.FindOne was just called")
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
//	len(mockedPostService.FindOneCalls())
func (mock *PostServiceMock) FindOneCalls() []struct {
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
