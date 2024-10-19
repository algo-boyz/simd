//go:build amd64
// +build amd64

package avx2

func Supported() bool

func AddInt32(left, right, result []int32) int

func AddInt64(left, right, result []int64) int

func AndInt32(left, right, result []int32) int

func AndInt64(left, right, result []int64) int

func MulInt32(left, right, result []int32) int

func OrInt32(left, right, result []int32) int

func OrInt64(left, right, result []int64) int

func SubInt32(left, right, result []int32) int

func SubInt64(left, right, result []int64) int

func XorInt32(left, right, result []int32) int

func XorInt64(left, right, result []int64) int
