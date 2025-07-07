package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// go test -v -run [nama tes],untuk menjalankan test yang di inginkan
//go test -list . ,untuk cek list test yang ada

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

// subtest
// go test -v -run TestSubtestAdd
// go test -v -run 'TestSubTestAdd/AddPertama',menjalankan sesuai dengan nama subtest yang di inginkan
func TestSubTestAdd(t *testing.T) {
	t.Run("AddPertama", func(t *testing.T) {
		result := Add(5, 7)
		require.Equal(t, 12, result, "5 + 7 harusnya 12")
	})
	t.Run("AddKedua", func(t *testing.T) {
		result := Add(8, 7)
		require.Equal(t, 15, result, "8 + 7 harusnya 15")
	})
	t.Run("AddKetiga", func(t *testing.T) {
		result := Add(10, 9)
		require.Equal(t, 19, result, "10 + 9 harusnya 19")
	})
}

// meneraptkan table test
func TestAddTable(t *testing.T) {
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
