// Simd support for arithmetic operations. Allowing for parallel element-wise computations.
package simd

import (
	"github.com/pehringer/simd/internal/fallback"
	"github.com/pehringer/simd/internal/shared"
)

var (
	addFloat32 shared.Operation[float32] = fallback.Add[float32]
	addFloat64 shared.Operation[float64] = fallback.Add[float64]
	addInt32   shared.Operation[int32]   = fallback.Add[int32]
	addInt64   shared.Operation[int64]   = fallback.Add[int64]
	andInt32   shared.Operation[int32]   = fallback.And[int32]
	andInt64   shared.Operation[int64]   = fallback.And[int64]
	divFloat32 shared.Operation[float32] = fallback.Div[float32]
	divFloat64 shared.Operation[float64] = fallback.Div[float64]
	divInt32   shared.Operation[int32]   = fallback.Div[int32]
	divInt64   shared.Operation[int64]   = fallback.Div[int64]
	mulFloat32 shared.Operation[float32] = fallback.Mul[float32]
	mulFloat64 shared.Operation[float64] = fallback.Mul[float64]
	mulInt32   shared.Operation[int32]   = fallback.Mul[int32]
	mulInt64   shared.Operation[int64]   = fallback.Mul[int64]
	orInt32    shared.Operation[int32]   = fallback.Or[int32]
	orInt64    shared.Operation[int64]   = fallback.Or[int64]
	subFloat32 shared.Operation[float32] = fallback.Sub[float32]
	subFloat64 shared.Operation[float64] = fallback.Sub[float64]
	subInt32   shared.Operation[int32]   = fallback.Sub[int32]
	subInt64   shared.Operation[int64]   = fallback.Sub[int64]
	xorInt32   shared.Operation[int32]   = fallback.Xor[int32]
	xorInt64   shared.Operation[int64]   = fallback.Xor[int64]
)

// AddFloat32 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddFloat32(left, right, result []float32) int {
	return addFloat32(left, right, result)
}

// AddFloat64 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddFloat64(left, right, result []float64) int {
	return addFloat64(left, right, result)
}

// AddInt32 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddInt32(left, right, result []int32) int {
	return addInt32(left, right, result)
}

// AddInt64 performs element-wise addition on left and right, storing the sums in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AddInt64(left, right, result []int64) int {
	return addInt64(left, right, result)
}

// AndInt32 performs element-wise AND on left and right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AndInt32(left, right, result []int32) int {
	return andInt32(left, right, result)
}

// AndInt64 performs element-wise AND on left and right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func AndInt64(left, right, result []int64) int {
	return andInt64(left, right, result)
}

// DivFloat32 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivFloat32(left, right, result []float32) int {
	return divFloat32(left, right, result)
}

// DivFloat64 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivFloat64(left, right, result []float64) int {
	return divFloat64(left, right, result)
}

// DivInt32 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivInt32(left, right, result []int32) int {
	return divInt32(left, right, result)
}

// DivInt64 performs element-wise division on left and right, storing the quotients in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func DivInt64(left, right, result []int64) int {
	return divInt64(left, right, result)
}

// MulFloat32 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulFloat32(left, right, result []float32) int {
	return mulFloat32(left, right, result)
}

// MulFloat64 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulFloat64(left, right, result []float64) int {
	return mulFloat64(left, right, result)
}

// MulInt32 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulInt32(left, right, result []int32) int {
	return mulInt32(left, right, result)
}

// MulInt64 performs element-wise multiplication on left and right, storing the products in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func MulInt64(left, right, result []int64) int {
	return mulInt64(left, right, result)
}

// OrInt32 performs element-wise OR on left and right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func OrInt32(left, right, result []int32) int {
	return orInt32(left, right, result)
}

// OrInt64 performs element-wise OR on left and right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func OrInt64(left, right, result []int64) int {
	return orInt64(left, right, result)
}

// SubFloat32 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubFloat32(left, right, result []float32) int {
	return subFloat32(left, right, result)
}

// SubFloat64 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubFloat64(left, right, result []float64) int {
	return subFloat64(left, right, result)
}

// SubInt32 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubInt32(left, right, result []int32) int {
	return subInt32(left, right, result)
}

// SubInt64 performs element-wise subtraction on left and right, storing the differences in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func SubInt64(left, right, result []int64) int {
	return subInt64(left, right, result)
}

// XorInt32 performs element-wise XOR on left and right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func XorInt32(left, right, result []int32) int {
	return xorInt32(left, right, result)
}

// XorInt64 performs element-wise XOR on left and right, storing the results in result.
// The operation is performed up to the shortest length of left, right, and result.
// Returns the number of operations performed.
func XorInt64(left, right, result []int64) int {
	return xorInt64(left, right, result)
}
