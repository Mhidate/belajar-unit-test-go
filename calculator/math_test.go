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

// meneraptkan table test
func TestAddDua(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positif case", 3, 2, 5},
		{"nol case", 0, 0, 0},
		{"negatif case", -2, -3, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}
