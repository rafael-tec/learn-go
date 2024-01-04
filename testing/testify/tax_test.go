package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)
}

func TestErrorCalculateTax(t *testing.T) {
	tax, err := CalculateTax(0)

	assert.Equal(t, 0.0, tax)
	assert.Error(t, err, "amount must be greater than 0")
}
