// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/rzfhlv/doit/modules/investor/model"
	mock "github.com/stretchr/testify/mock"

	param "github.com/rzfhlv/doit/utilities/param"
)

// IUsecase is an autogenerated mock type for the IUsecase type
type IUsecase struct {
	mock.Mock
}

// ConventionalMigrate provides a mock function with given fields: ctx
func (_m *IUsecase) ConventionalMigrate(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Generate provides a mock function with given fields: ctx
func (_m *IUsecase) Generate(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, _a1
func (_m *IUsecase) GetAll(ctx context.Context, _a1 *param.Param) ([]model.Investor, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []model.Investor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *param.Param) ([]model.Investor, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *param.Param) []model.Investor); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Investor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *param.Param) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IUsecase) GetByID(ctx context.Context, id int64) (model.Investor, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Investor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.Investor, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Investor); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Investor)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MigrateInvestors provides a mock function with given fields: ctx
func (_m *IUsecase) MigrateInvestors(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUsecase creates a new instance of IUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUsecase {
	mock := &IUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}