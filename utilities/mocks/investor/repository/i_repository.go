// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/rzfhlv/doit/modules/investor/model"
	mock "github.com/stretchr/testify/mock"

	param "github.com/rzfhlv/doit/utilities/param"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *IRepository) Count(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOutbox provides a mock function with given fields: ctx, identifier
func (_m *IRepository) DeleteOutbox(ctx context.Context, identifier int64) error {
	ret := _m.Called(ctx, identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Generate provides a mock function with given fields: ctx, name
func (_m *IRepository) Generate(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, _a1
func (_m *IRepository) GetAll(ctx context.Context, _a1 param.Param) ([]model.Investor, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []model.Investor
	if rf, ok := ret.Get(0).(func(context.Context, param.Param) []model.Investor); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Investor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, param.Param) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IRepository) GetByID(ctx context.Context, id int64) (model.Investor, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Investor
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Investor); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Investor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPsql provides a mock function with given fields: ctx
func (_m *IRepository) GetPsql(ctx context.Context) ([]model.Investor, error) {
	ret := _m.Called(ctx)

	var r0 []model.Investor
	if rf, ok := ret.Get(0).(func(context.Context) []model.Investor); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Investor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMongo provides a mock function with given fields: ctx, investor
func (_m *IRepository) SaveMongo(ctx context.Context, investor model.Investor) error {
	ret := _m.Called(ctx, investor)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Investor) error); ok {
		r0 = rf(ctx, investor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertMongo provides a mock function with given fields: ctx, investor
func (_m *IRepository) UpsertMongo(ctx context.Context, investor model.Investor) error {
	ret := _m.Called(ctx, investor)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Investor) error); ok {
		r0 = rf(ctx, investor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertOutbox provides a mock function with given fields: ctx, outbox
func (_m *IRepository) UpsertOutbox(ctx context.Context, outbox model.Outbox) error {
	ret := _m.Called(ctx, outbox)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Outbox) error); ok {
		r0 = rf(ctx, outbox)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepository(t mockConstructorTestingTNewIRepository) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}