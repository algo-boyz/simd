package simd

import (
	"testing"
	"math/rand"
	"time"
)

// Helper function to create random float32 slices
func createRandomSlice(size int) []float32 {
	rand.Seed(time.Now().UnixNano())
	slice := make([]float32, size)
	for i := range slice {
		slice[i] = rand.Float32()
	}
	return slice
}

// Benchmark for Function1
func BenchmarkCode(b *testing.B) {
	size := 4096 // Choose the slice size
	a := createRandomSlice(size)
	bb := createRandomSlice(size)
	result := make([]float32, size)

	// Reset the timer and run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(a, bb, result)
	}
}

// Benchmark for Function2
func BenchmarkSimd(b *testing.B) {
	size := 4096 // Choose the slice size
	a := createRandomSlice(size)
	bb := createRandomSlice(size)
	result := make([]float32, size)

	// Reset the timer and run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sseAddFloat32(a, bb, result)
	}
}

func TestSimd(t *testing.T) {
	// Define a small prime number for the array size
	size := 11

	// Create test data
	a := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 0}
	b := []float32{10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 1000.0, 1100.0, 0, 0}
	expected := []float32{11.0, 22.0, 33.0, 44.0, 55.0, 66.0, 77.0, 88.0, 99.0, 1010.0, 1111.0}

	// Result slices
	result := make([]float32, size)

	// Call both functions
	if n := sseAddFloat32(a, b, result); n != size {
		t.Errorf("Simd failed returned %v, want %v", n, size)
        }
	// Check if results match expected values
	for i := 0; i < size; i++ {
		if result[i] != expected[i] {
			t.Errorf("failed failed at index %d: got %v, want %v", i, result[i], expected[i])
		}
	}
}
