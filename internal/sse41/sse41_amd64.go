//go:build amd64
// +build amd64

package sse41

func Supported() bool

func MulInt32(left, right, result []int32) int
