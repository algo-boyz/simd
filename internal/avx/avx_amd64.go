//go:build amd64
// +build amd64

package avx

func Supported() bool

func AddFloat32(left, right, result []float32) int

func AddFloat64(left, right, result []float64) int

func DivFloat32(left, right, result []float32) int

func DivFloat64(left, right, result []float64) int

func MulFloat32(left, right, result []float32) int

func MulFloat64(left, right, result []float64) int

func SubFloat32(left, right, result []float32) int

func SubFloat64(left, right, result []float64) int
