package helper

import (
	"github.com/stretchr/testify/mock"
)

type Where interface {
	From() Field
	Operator() string
	To() any

	SetFrom(from Field) Where
	SetOperator(operator string) Where
	SetTo(to any) Where
}

// FakeWhere is an autogenerated mock type for the FakeWhere type
type FakeWhere struct {
	mock.Mock
}

// From provides a mock function with given fields:
func (_m *FakeWhere) From() Field {
	ret := _m.Called()

	var r0 Field
	if rf, ok := ret.Get(0).(func() Field); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Field)
		}
	}

	return r0
}

// Operator provides a mock function with given fields:
func (_m *FakeWhere) Operator() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetFrom provides a mock function with given fields: from
func (_m *FakeWhere) SetFrom(from Field) Where {
	ret := _m.Called(from)

	var r0 Where
	if rf, ok := ret.Get(0).(func(Field) Where); ok {
		r0 = rf(from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Where)
		}
	}

	return r0
}

// SetOperator provides a mock function with given fields: operator
func (_m *FakeWhere) SetOperator(operator string) Where {
	ret := _m.Called(operator)

	var r0 Where
	if rf, ok := ret.Get(0).(func(string) Where); ok {
		r0 = rf(operator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Where)
		}
	}

	return r0
}

// SetTo provides a mock function with given fields: to
func (_m *FakeWhere) SetTo(to interface{}) Where {
	ret := _m.Called(to)

	var r0 Where
	if rf, ok := ret.Get(0).(func(interface{}) Where); ok {
		r0 = rf(to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Where)
		}
	}

	return r0
}

// To provides a mock function with given fields:
func (_m *FakeWhere) To() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type mockConstructorTestingTNewWhere interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeWhere creates a new instance of FakeWhere. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeWhere(t mockConstructorTestingTNewWhere) *FakeWhere {
	mock := &FakeWhere{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}