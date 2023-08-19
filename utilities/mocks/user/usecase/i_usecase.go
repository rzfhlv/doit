// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	jwt "github.com/rzfhlv/doit/utilities/jwt"
	mock "github.com/stretchr/testify/mock"

	model "github.com/rzfhlv/doit/modules/user/model"
)

// IUsecase is an autogenerated mock type for the IUsecase type
type IUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, login
func (_m *IUsecase) Login(ctx context.Context, login model.Login) (model.JWT, error) {
	ret := _m.Called(ctx, login)

	var r0 model.JWT
	if rf, ok := ret.Get(0).(func(context.Context, model.Login) model.JWT); ok {
		r0 = rf(ctx, login)
	} else {
		r0 = ret.Get(0).(model.JWT)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Login) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, token
func (_m *IUsecase) Logout(ctx context.Context, token string) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: ctx, user
func (_m *IUsecase) Register(ctx context.Context, user model.User) (model.JWT, error) {
	ret := _m.Called(ctx, user)

	var r0 model.JWT
	if rf, ok := ret.Get(0).(func(context.Context, model.User) model.JWT); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(model.JWT)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: ctx, validate
func (_m *IUsecase) Validate(ctx context.Context, validate model.Validate) (*jwt.JWTClaim, error) {
	ret := _m.Called(ctx, validate)

	var r0 *jwt.JWTClaim
	if rf, ok := ret.Get(0).(func(context.Context, model.Validate) *jwt.JWTClaim); ok {
		r0 = rf(ctx, validate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.JWTClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Validate) error); ok {
		r1 = rf(ctx, validate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUsecase creates a new instance of IUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUsecase(t mockConstructorTestingTNewIUsecase) *IUsecase {
	mock := &IUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}