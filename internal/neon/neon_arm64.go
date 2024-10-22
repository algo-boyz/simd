//go:build arm64
// +build arm64

package neon

// NEON Vector Instructions used:
// VLD1.4S - Load 4 single-precision floats into a vector register
// VST1.4S - Store 4 single-precision floats from a vector register
// FADD V1.4S, V0.4S, V2 - Add 4 floats in parallel
func AddFloat32(left, right, result []float32) int

func AddInt64(x, y int64) int64

func MulFloat64(left, right, result []float64) int
