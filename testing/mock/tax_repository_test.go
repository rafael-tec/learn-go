package mock_test

import (
	"errors"
	"mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveTax(t *testing.T) {
	dbClientMock := &mock.DatabaseClientMock{}
	repository := mock.NewTaxRepository(dbClientMock)

	expectedTax := 10.0

	dbClientMock.
		On("Save", "", mock.Entity{ID: "a2b4c0g9", Tax: expectedTax}).
		Return(nil)

	err := repository.SaveTax(expectedTax)

	assert.Nil(t, err)

	dbClientMock.AssertNumberOfCalls(t, "Save", 1)
}

func TestErrorSaveTax(t *testing.T) {
	dbClientMock := &mock.DatabaseClientMock{}
	repository := mock.NewTaxRepository(dbClientMock)

	expectedTax := 10.0
	expectedMsgErr := "unexpected error"

	dbClientMock.
		On("Save", "", mock.Entity{ID: "a2b4c0g9", Tax: expectedTax}).
		Return(errors.New(expectedMsgErr))

	err := repository.SaveTax(expectedTax)

	assert.EqualError(t, err, expectedMsgErr)

	dbClientMock.AssertExpectations(t)
}
