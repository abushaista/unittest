package unittest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSalah(t *testing.T) {
	result := Calculate(8)
	t.Log(result)

	require.Equal(t, 4.0, result)
	assert.Equal(t, 4.0, result)
	fmt.Println("nilai salah")
}

func TestCalculateFailed(t *testing.T) {
	result := Calculate(10)
	fmt.Println(result)
	if result != 5.0 {
		t.Fatal("nilai salah")
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(i)
	}
}
