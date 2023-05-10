package mocks

import (
	"github.com/kitstack/dbkit/specs"
	"github.com/stretchr/testify/mock"
)

// FakeConnectorsInstance is an autogenerated mock type for the FakeConnectorsInstance type
type FakeConnectorsInstance struct {
	mock.Mock
}

// Instance provides a mock function with given fields:
func (_m *FakeConnectorsInstance) Instance() specs.Connectors {
	ret := _m.Called()

	var r0 specs.Connectors
	if rf, ok := ret.Get(0).(func() specs.Connectors); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(specs.Connectors)
		}
	}

	return r0
}

type mockConstructorTestingTNewConnectorsInstance interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeConnectorsInstance creates a new instance of FakeConnectorsInstance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeConnectorsInstance(t mockConstructorTestingTNewConnectorsInstance) *FakeConnectorsInstance {
	fakeConnectorsInstance := &FakeConnectorsInstance{}
	fakeConnectorsInstance.Mock.Test(t)

	t.Cleanup(func() { fakeConnectorsInstance.AssertExpectations(t) })

	return fakeConnectorsInstance
}