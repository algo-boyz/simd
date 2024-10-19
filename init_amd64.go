//go:build amd64
// +build amd64

package simd

import (
	"github.com/pehringer/simd/internal/avx"
	"github.com/pehringer/simd/internal/avx2"
	"github.com/pehringer/simd/internal/sse2"
	"github.com/pehringer/simd/internal/sse41"
)

func init() {
	// if sse.Supported() {
	// 	addFloat32 = sse.AddFloat32
	// 	divFloat32 = sse.DivFloat32
	// 	mulFloat32 = sse.MulFloat32
	// 	subFloat32 = sse.SubFloat32
	// }
	if sse2.Supported() {
		addFloat64 = sse2.AddFloat64
		addInt32 = sse2.AddInt32
		addInt64 = sse2.AddInt64
		andInt32 = sse2.AndInt32
		andInt64 = sse2.AndInt64
		divFloat64 = sse2.DivFloat64
		mulFloat64 = sse2.MulFloat64
		orInt32 = sse2.OrInt32
		orInt64 = sse2.OrInt64
		subFloat64 = sse2.SubFloat64
		subInt32 = sse2.SubInt32
		subInt64 = sse2.SubInt64
		xorInt32 = sse2.XorInt32
		xorInt64 = sse2.XorInt64
	}
	if sse41.Supported() {
		mulInt32 = sse41.MulInt32
	}
	if avx.Supported() {
		addFloat32 = avx.AddFloat32
		addFloat64 = avx.AddFloat64
		divFloat32 = avx.DivFloat32
		divFloat64 = avx.DivFloat64
		mulFloat32 = avx.MulFloat32
		mulFloat64 = avx.MulFloat64
		subFloat32 = avx.SubFloat32
		subFloat64 = avx.SubFloat64
	}
	if avx2.Supported() {
		addInt32 = avx2.AddInt32
		addInt64 = avx2.AddInt64
		andInt32 = avx2.AndInt32
		andInt64 = avx2.AndInt64
		mulInt32 = avx2.MulInt32
		orInt32 = avx2.OrInt32
		orInt64 = avx2.OrInt64
		subInt32 = avx2.SubInt32
		subInt64 = sse2.SubInt64
		xorInt32 = avx2.XorInt32
		xorInt64 = avx2.XorInt64
	}
}
