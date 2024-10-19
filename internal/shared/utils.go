package shared

import (
	"math"
	"testing"
)

// CheckSlice for same length, capacity, and element equality.
func CheckSlice[T Floating | Integer](t *testing.T, test, control []T) bool {
	if len(test) != len(control) {
		t.Errorf("lengths not equal")
		return false
	}
	if cap(test) != cap(control) {
		t.Errorf("capacities not equal")
		return false
	}
	for i := 0; i < len(control); i++ {
		if test[i] != control[i] {
			t.Errorf("elements not equal")
			return false
		}
	}
	return true
}

// CheckOperation checks for correct length and that the input slices are unchanged.
// Returns true if the operation is correct, false otherwise.
func CheckOperation[T Floating | Integer](t *testing.T, test, control Operation[T], left, right, result []T) bool {
	testLeft := make([]T, len(left), cap(left))
	copy(testLeft, left)
	testRight := make([]T, len(right), cap(right))
	copy(testRight, right)
	testResult := make([]T, len(result), cap(result))
	copy(testResult, result)
	if test(testLeft, testRight, testResult) != control(left, right, result) {
		t.Errorf("operation returned incorrect length")
		return false
	}
	if !CheckSlice(t, testLeft, left) {
		return false
	}
	if !CheckSlice(t, testRight, right) {
		return false
	}
	if !CheckSlice(t, testResult, result) {
		return false
	}
	return true
}

var (
	increment int32 = 1
	decrement int32 = math.MaxInt32
)

// Large returns a slice of length elements, starting at math.MaxInt32 and decrementing by 1.
func Large[T Floating | Integer](length int) []T {
	elements := make([]T, length)
	for i := 0; i < length; i++ {
		elements[i] = T(decrement)
		decrement--
	}
	return elements
}

// Small returns a slice of length elements, starting at 1 and incrementing by 1.
func Small[T Floating | Integer](length int) []T {
	elements := make([]T, length)
	for i := 0; i < length; i++ {
		elements[i] = T(increment)
		increment++
	}
	return elements
}
