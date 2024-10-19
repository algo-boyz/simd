package fallback

import (
	"fmt"

	"github.com/pehringer/simd/internal/shared"
)

// Add performs element-wise addition on left to right, storing sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func Add[T shared.Floating | shared.Integer](left, right, result []T) int {
	fmt.Println("fallback Add")
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] + right[i]
		result[i+1] = left[i+1] + right[i+1]
		result[i+2] = left[i+2] + right[i+2]
		result[i+3] = left[i+3] + right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] + right[i]
	}
	return n
}

// And performs element-wise bitwise AND on left to right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func And[T shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] & right[i]
		result[i+1] = left[i+1] & right[i+1]
		result[i+2] = left[i+2] & right[i+2]
		result[i+3] = left[i+3] & right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] & right[i]
	}
	return n
}

// Div performs element-wise division on left to right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func Div[T shared.Floating | shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] / right[i]
		result[i+1] = left[i+1] / right[i+1]
		result[i+2] = left[i+2] / right[i+2]
		result[i+3] = left[i+3] / right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] / right[i]
	}
	return n
}

// Mul performs element-wise multiplication on left to right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func Mul[T shared.Floating | shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] * right[i]
		result[i+1] = left[i+1] * right[i+1]
		result[i+2] = left[i+2] * right[i+2]
		result[i+3] = left[i+3] * right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] * right[i]
	}
	return n
}

// Or performs element-wise bitwise OR on left to right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func Or[T shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] | right[i]
		result[i+1] = left[i+1] | right[i+1]
		result[i+2] = left[i+2] | right[i+2]
		result[i+3] = left[i+3] | right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] | right[i]
	}
	return n
}

// Sub performs element-wise subtraction on left to right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func Sub[T shared.Floating | shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] - right[i]
		result[i+1] = left[i+1] - right[i+1]
		result[i+2] = left[i+2] - right[i+2]
		result[i+3] = left[i+3] - right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] - right[i]
	}
	return n
}

// Xor performs element-wise bitwise XOR on left to right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func Xor[T shared.Integer](left, right, result []T) int {
	n := len(result)
	if m := len(left); m < n {
		n = m
	}
	if m := len(right); m < n {
		n = m
	}
	i := 0
	for ; n-i >= 4; i += 4 {
		result[i] = left[i] ^ right[i]
		result[i+1] = left[i+1] ^ right[i+1]
		result[i+2] = left[i+2] ^ right[i+2]
		result[i+3] = left[i+3] ^ right[i+3]
	}
	for ; i < n; i++ {
		result[i] = left[i] ^ right[i]
	}
	return n
}
