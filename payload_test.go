package dbkit

import (
	"context"
	"errors"
	"github.com/lab210-dev/dbkit/specs"
	"github.com/lab210-dev/dbkit/tests/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PayloadTestSuite struct {
	suite.Suite
	context.Context
	fakeModelDefinition *mocks.FakeModelDefinition
	fakeFieldDefinition *mocks.FakeFieldDefinition
	fakeDriverField     *mocks.FakeDriverField
}

func (test *PayloadTestSuite) SetupTest() {
	test.Context = context.Background()
	test.fakeModelDefinition = mocks.NewFakeModelDefinition(test.T())
	test.fakeFieldDefinition = mocks.NewFakeFieldDefinition(test.T())
	test.fakeDriverField = mocks.NewFakeDriverField(test.T())
}

func (test *PayloadTestSuite) TestMappingWithGetFieldByNameErr() {
	newPayload := NewPayload[specs.Model]()
	tmp := newPayload.(*payload[specs.Model])

	tmp.modelDefinition = test.fakeModelDefinition

	newPayload.SetFields([]specs.DriverField{
		test.fakeDriverField,
	})

	test.fakeDriverField.On("NameInSchema").Return("Test").Once()
	test.fakeModelDefinition.On("GetFieldByName", "Test").Return(nil, errors.New("test")).Once()

	_, err := newPayload.Mapping()
	test.Error(err)
}

func (test *PayloadTestSuite) TestMappingSuccessful() {
	newPayload := NewPayload[specs.Model]()
	tmp := newPayload.(*payload[specs.Model])

	tmp.modelDefinition = test.fakeModelDefinition

	newPayload.SetFields([]specs.DriverField{
		test.fakeDriverField,
	})

	test.fakeDriverField.On("NameInSchema").Return("Test").Once()
	test.fakeModelDefinition.On("GetFieldByName", "Test").Return(test.fakeFieldDefinition, nil).Once()

	test.fakeFieldDefinition.On("Copy").Return(nil).Once()

	values, err := newPayload.Mapping()
	if !test.NoError(err) {
		return
	}
	test.Len(values, 1)
}

func (test *PayloadTestSuite) TestOnScanGetFieldByNameErr() {
	newPayload := NewPayload[specs.Model]()
	tmp := newPayload.(*payload[specs.Model])

	tmp.modelDefinition = test.fakeModelDefinition

	newPayload.SetFields([]specs.DriverField{
		test.fakeDriverField,
	})

	test.fakeDriverField.On("NameInSchema").Return("Test").Once()
	test.fakeModelDefinition.On("GetFieldByName", "Test").Return(nil, errors.New("test")).Once()

	err := newPayload.OnScan([]any{})
	test.Error(err)
}

func (test *PayloadTestSuite) TestJoin() {
	newPayload := NewPayload[specs.Model]()
	tmp := newPayload.(*payload[specs.Model])

	var t []specs.DriverJoin
	test.Equal(tmp.Join(), t)
}

func TestPayloadTestSuite(t *testing.T) {
	suite.Run(t, new(PayloadTestSuite))
}