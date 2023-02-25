// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package fakesql

import (
	"database/sql/driver"
	mock "github.com/stretchr/testify/mock"
)

// FakeStmt is an autogenerated mock type for the FakeStmt type
type FakeStmt struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *FakeStmt) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: args
func (_m *FakeStmt) Exec(args []driver.Value) (driver.Result, error) {
	ret := _m.Called(args)

	var r0 driver.Result
	if rf, ok := ret.Get(0).(func([]driver.Value) driver.Result); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]driver.Value) error); ok {
		r1 = rf(args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NumInput provides a mock function with given fields:
func (_m *FakeStmt) NumInput() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Query provides a mock function with given fields: args
func (_m *FakeStmt) Query(args []driver.Value) (driver.Rows, error) {
	ret := _m.Called(args)

	var r0 driver.Rows
	if rf, ok := ret.Get(0).(func([]driver.Value) driver.Rows); ok {
		r0 = rf(args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]driver.Value) error); ok {
		r1 = rf(args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStmt interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeStmt creates a new instance of FakeStmt. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeStmt(t mockConstructorTestingTNewStmt) *FakeStmt {
	mock := &FakeStmt{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}