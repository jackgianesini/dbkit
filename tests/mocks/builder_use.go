package mocks

import (
	"context"

	"github.com/kitstack/dbkit/specs"
	"github.com/stretchr/testify/mock"
)

// FakeBuilderUse is an autogenerated mock type for the FakeBuilderUse type
type FakeBuilderUse[T specs.Model] struct {
	mock.Mock
}

// Use provides a mock function with given fields: ctx, connector
func (_m *FakeBuilderUse[T]) Use(ctx context.Context, connector specs.Connector) specs.Builder[T] {
	ret := _m.Called(ctx, connector)

	var r0 specs.Builder[T]
	if rf, ok := ret.Get(0).(func(context.Context, specs.Connector) specs.Builder[T]); ok {
		r0 = rf(ctx, connector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(specs.Builder[T])
		}
	}

	return r0
}

type mockConstructorTestingTNewBuilderUse interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeBuilderUse creates a new instance of FakeBuilderUse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeBuilderUse[T specs.Model](t mockConstructorTestingTNewBuilderUse) *FakeBuilderUse[T] {
	fakeBuilderUse := &FakeBuilderUse[T]{}
	fakeBuilderUse.Mock.Test(t)

	t.Cleanup(func() { fakeBuilderUse.AssertExpectations(t) })

	return fakeBuilderUse
}
