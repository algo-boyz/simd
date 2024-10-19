//go:build amd64
// +build amd64

package sse2

func Supported() bool

func AddFloat64(left, right, result []float64) int

func AddInt32(left, right, result []int32) int

func AddInt64(left, right, result []int64) int

func AndInt32(left, right, result []int32) int

func AndInt64(left, right, result []int64) int

func DivFloat64(left, right, result []float64) int

func MulFloat64(left, right, result []float64) int

func OrInt32(left, right, result []int32) int

func OrInt64(left, right, result []int64) int

func SubFloat64(left, right, result []float64) int

func SubInt32(left, right, result []int32) int

func SubInt64(left, right, result []int64) int

func XorInt32(left, right, result []int32) int

func XorInt64(left, right, result []int64) int
