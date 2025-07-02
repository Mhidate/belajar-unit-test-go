package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	assert.Equal(t, 7, result, "3 + 4 harusnya 7")
}

func TestDivide(t *testing.T) {
	// Test pembagian normal
	result, err := Divide(10, 2)
	assert.NoError(t, err)
	assert.Equal(t, 5, result)

	// Test divide by zero
	result, err = Divide(10, 0)
	assert.Error(t, err)
	assert.Equal(t, 0, result)
}
