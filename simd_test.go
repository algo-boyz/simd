package simd

import (
	"fmt"
	"testing"

	"github.com/pehringer/simd/internal/fallback"
	"github.com/pehringer/simd/internal/shared"
	"github.com/stretchr/testify/assert"
)

func TestAddFloat32Zero(t *testing.T) {
	shared.CheckOperation(t, AddFloat32, fallback.Add, []float32{}, []float32{}, []float32{})
}

func TestAddFloat32Prime(t *testing.T) {
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Small[float32](11), shared.Small[float32](13), shared.Small[float32](17))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Small[float32](29), shared.Small[float32](19), shared.Large[float32](23))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Small[float32](37), shared.Large[float32](41), shared.Small[float32](31))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Small[float32](43), shared.Large[float32](47), shared.Large[float32](53))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Large[float32](67), shared.Small[float32](59), shared.Small[float32](61))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Large[float32](73), shared.Small[float32](79), shared.Large[float32](71))
}

func TestAddFloat32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Large[float32](1024), shared.Large[float32](1024), shared.Small[float32](1024))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Large[float32](2048), shared.Large[float32](2048), shared.Large[float32](2048))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Small[float32](4096), shared.Small[float32](4096), shared.Small[float32](4096))
	shared.CheckOperation(t, AddFloat32, fallback.Add, shared.Small[float32](8192), shared.Small[float32](8192), shared.Large[float32](8192))
}

func TestAddFloat64Zero(t *testing.T) {
	shared.CheckOperation(t, AddFloat64, fallback.Add, []float64{}, []float64{}, []float64{})
}

func TestAddFloat64Prime(t *testing.T) {
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Small[float64](11), shared.Small[float64](13), shared.Small[float64](17))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Small[float64](29), shared.Small[float64](19), shared.Large[float64](23))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Small[float64](37), shared.Large[float64](41), shared.Small[float64](31))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Small[float64](43), shared.Large[float64](47), shared.Large[float64](53))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Large[float64](67), shared.Small[float64](59), shared.Small[float64](61))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Large[float64](73), shared.Small[float64](79), shared.Large[float64](71))
}

func TestAddFloat64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Large[float64](1024), shared.Large[float64](1024), shared.Small[float64](1024))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Large[float64](2048), shared.Large[float64](2048), shared.Large[float64](2048))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Small[float64](4096), shared.Small[float64](4096), shared.Small[float64](4096))
	shared.CheckOperation(t, AddFloat64, fallback.Add, shared.Small[float64](8192), shared.Small[float64](8192), shared.Large[float64](8192))
}

func TestAddInt32Zero(t *testing.T) {
	shared.CheckOperation(t, AddInt32, fallback.Add, []int32{}, []int32{}, []int32{})
}

func TestAddInt32Prime(t *testing.T) {
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Small[int32](11), shared.Small[int32](13), shared.Small[int32](17))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Small[int32](29), shared.Small[int32](19), shared.Large[int32](23))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Small[int32](37), shared.Large[int32](41), shared.Small[int32](31))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Small[int32](43), shared.Large[int32](47), shared.Large[int32](53))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Large[int32](67), shared.Small[int32](59), shared.Small[int32](61))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Large[int32](73), shared.Small[int32](79), shared.Large[int32](71))
}

func TestAddInt32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Large[int32](1024), shared.Large[int32](1024), shared.Small[int32](1024))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Large[int32](2048), shared.Large[int32](2048), shared.Large[int32](2048))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Small[int32](4096), shared.Small[int32](4096), shared.Small[int32](4096))
	shared.CheckOperation(t, AddInt32, fallback.Add, shared.Small[int32](8192), shared.Small[int32](8192), shared.Large[int32](8192))
}

func TestAddInt64Zero(t *testing.T) {
	shared.CheckOperation(t, AddInt64, fallback.Add, []int64{}, []int64{}, []int64{})
}

func TestAddInt64Prime(t *testing.T) {
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Small[int64](11), shared.Small[int64](13), shared.Small[int64](17))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Small[int64](29), shared.Small[int64](19), shared.Large[int64](23))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Small[int64](37), shared.Large[int64](41), shared.Small[int64](31))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Small[int64](43), shared.Large[int64](47), shared.Large[int64](53))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Large[int64](67), shared.Small[int64](59), shared.Small[int64](61))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Large[int64](73), shared.Small[int64](79), shared.Large[int64](71))
}

func TestAddInt64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Large[int64](1024), shared.Large[int64](1024), shared.Small[int64](1024))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Large[int64](2048), shared.Large[int64](2048), shared.Large[int64](2048))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Small[int64](4096), shared.Small[int64](4096), shared.Small[int64](4096))
	shared.CheckOperation(t, AddInt64, fallback.Add, shared.Small[int64](8192), shared.Small[int64](8192), shared.Large[int64](8192))
}

func TestAndInt32Zero(t *testing.T) {
	shared.CheckOperation(t, AndInt32, fallback.And, []int32{}, []int32{}, []int32{})
}

func TestAndInt32Prime(t *testing.T) {
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Small[int32](11), shared.Small[int32](13), shared.Small[int32](17))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Small[int32](29), shared.Small[int32](19), shared.Large[int32](23))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Small[int32](37), shared.Large[int32](41), shared.Small[int32](31))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Small[int32](43), shared.Large[int32](47), shared.Large[int32](53))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Large[int32](67), shared.Small[int32](59), shared.Small[int32](61))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Large[int32](73), shared.Small[int32](79), shared.Large[int32](71))
}

func TestAndInt32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Large[int32](1024), shared.Large[int32](1024), shared.Small[int32](1024))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Large[int32](2048), shared.Large[int32](2048), shared.Large[int32](2048))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Small[int32](4096), shared.Small[int32](4096), shared.Small[int32](4096))
	shared.CheckOperation(t, AndInt32, fallback.And, shared.Small[int32](8192), shared.Small[int32](8192), shared.Large[int32](8192))
}

func TestAndInt64Zero(t *testing.T) {
	shared.CheckOperation(t, AndInt64, fallback.And, []int64{}, []int64{}, []int64{})
}

func TestAndInt64Prime(t *testing.T) {
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Small[int64](11), shared.Small[int64](13), shared.Small[int64](17))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Small[int64](29), shared.Small[int64](19), shared.Large[int64](23))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Small[int64](37), shared.Large[int64](41), shared.Small[int64](31))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Small[int64](43), shared.Large[int64](47), shared.Large[int64](53))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Large[int64](67), shared.Small[int64](59), shared.Small[int64](61))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Large[int64](73), shared.Small[int64](79), shared.Large[int64](71))
}

func TestAndInt64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Large[int64](1024), shared.Large[int64](1024), shared.Small[int64](1024))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Large[int64](2048), shared.Large[int64](2048), shared.Large[int64](2048))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Small[int64](4096), shared.Small[int64](4096), shared.Small[int64](4096))
	shared.CheckOperation(t, AndInt64, fallback.And, shared.Small[int64](8192), shared.Small[int64](8192), shared.Large[int64](8192))
}

func TestDivFloat32Zero(t *testing.T) {
	shared.CheckOperation(t, DivFloat32, fallback.Div, []float32{}, []float32{}, []float32{})
}

func TestDivFloat32Prime(t *testing.T) {
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Small[float32](11), shared.Small[float32](13), shared.Small[float32](17))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Small[float32](29), shared.Small[float32](19), shared.Large[float32](23))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Small[float32](37), shared.Large[float32](41), shared.Small[float32](31))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Small[float32](43), shared.Large[float32](47), shared.Large[float32](53))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Large[float32](67), shared.Small[float32](59), shared.Small[float32](61))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Large[float32](73), shared.Small[float32](79), shared.Large[float32](71))
}

func TestDivFloat32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Large[float32](1024), shared.Large[float32](1024), shared.Small[float32](1024))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Large[float32](2048), shared.Large[float32](2048), shared.Large[float32](2048))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Small[float32](4096), shared.Small[float32](4096), shared.Small[float32](4096))
	shared.CheckOperation(t, DivFloat32, fallback.Div, shared.Small[float32](8192), shared.Small[float32](8192), shared.Large[float32](8192))
}

func TestDivFloat64Zero(t *testing.T) {
	shared.CheckOperation(t, DivFloat64, fallback.Div, []float64{}, []float64{}, []float64{})
}

func TestDivFloat64Prime(t *testing.T) {
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Small[float64](11), shared.Small[float64](13), shared.Small[float64](17))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Small[float64](29), shared.Small[float64](19), shared.Large[float64](23))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Small[float64](37), shared.Large[float64](41), shared.Small[float64](31))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Small[float64](43), shared.Large[float64](47), shared.Large[float64](53))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Large[float64](67), shared.Small[float64](59), shared.Small[float64](61))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Large[float64](73), shared.Small[float64](79), shared.Large[float64](71))
}

func TestDivFloat64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Large[float64](1024), shared.Large[float64](1024), shared.Small[float64](1024))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Large[float64](2048), shared.Large[float64](2048), shared.Large[float64](2048))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Small[float64](4096), shared.Small[float64](4096), shared.Small[float64](4096))
	shared.CheckOperation(t, DivFloat64, fallback.Div, shared.Small[float64](8192), shared.Small[float64](8192), shared.Large[float64](8192))
}

func TestDivInt32Zero(t *testing.T) {
	shared.CheckOperation(t, DivInt32, fallback.Div, []int32{}, []int32{}, []int32{})
}

func TestDivInt32Prime(t *testing.T) {
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Small[int32](11), shared.Small[int32](13), shared.Small[int32](17))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Small[int32](29), shared.Small[int32](19), shared.Large[int32](23))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Small[int32](37), shared.Large[int32](41), shared.Small[int32](31))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Small[int32](43), shared.Large[int32](47), shared.Large[int32](53))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Large[int32](67), shared.Small[int32](59), shared.Small[int32](61))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Large[int32](73), shared.Small[int32](79), shared.Large[int32](71))
}

func TestDivInt32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Large[int32](1024), shared.Large[int32](1024), shared.Small[int32](1024))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Large[int32](2048), shared.Large[int32](2048), shared.Large[int32](2048))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Small[int32](4096), shared.Small[int32](4096), shared.Small[int32](4096))
	shared.CheckOperation(t, DivInt32, fallback.Div, shared.Small[int32](8192), shared.Small[int32](8192), shared.Large[int32](8192))
}

func TestDivInt64Zero(t *testing.T) {
	shared.CheckOperation(t, DivInt64, fallback.Div, []int64{}, []int64{}, []int64{})
}

func TestDivInt64Prime(t *testing.T) {
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Small[int64](11), shared.Small[int64](13), shared.Small[int64](17))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Small[int64](29), shared.Small[int64](19), shared.Large[int64](23))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Small[int64](37), shared.Large[int64](41), shared.Small[int64](31))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Small[int64](43), shared.Large[int64](47), shared.Large[int64](53))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Large[int64](67), shared.Small[int64](59), shared.Small[int64](61))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Large[int64](73), shared.Small[int64](79), shared.Large[int64](71))
}

func TestDivInt64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Large[int64](1024), shared.Large[int64](1024), shared.Small[int64](1024))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Large[int64](2048), shared.Large[int64](2048), shared.Large[int64](2048))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Small[int64](4096), shared.Small[int64](4096), shared.Small[int64](4096))
	shared.CheckOperation(t, DivInt64, fallback.Div, shared.Small[int64](8192), shared.Small[int64](8192), shared.Large[int64](8192))
}

func TestMulFloat32Zero(t *testing.T) {
	shared.CheckOperation(t, MulFloat32, fallback.Mul, []float32{}, []float32{}, []float32{})
}

func TestMulFloat32Prime(t *testing.T) {
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Small[float32](11), shared.Small[float32](13), shared.Small[float32](17))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Small[float32](29), shared.Small[float32](19), shared.Large[float32](23))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Small[float32](37), shared.Large[float32](41), shared.Small[float32](31))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Small[float32](43), shared.Large[float32](47), shared.Large[float32](53))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Large[float32](67), shared.Small[float32](59), shared.Small[float32](61))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Large[float32](73), shared.Small[float32](79), shared.Large[float32](71))
}

func TestMulFloat32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Large[float32](1024), shared.Large[float32](1024), shared.Small[float32](1024))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Large[float32](2048), shared.Large[float32](2048), shared.Large[float32](2048))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Small[float32](4096), shared.Small[float32](4096), shared.Small[float32](4096))
	shared.CheckOperation(t, MulFloat32, fallback.Mul, shared.Small[float32](8192), shared.Small[float32](8192), shared.Large[float32](8192))
}

func TestMulFloat64Zero(t *testing.T) {
	shared.CheckOperation(t, MulFloat64, fallback.Mul, []float64{}, []float64{}, []float64{})
}

func TestMulFloat64Prime(t *testing.T) {
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Small[float64](11), shared.Small[float64](13), shared.Small[float64](17))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Small[float64](29), shared.Small[float64](19), shared.Large[float64](23))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Small[float64](37), shared.Large[float64](41), shared.Small[float64](31))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Small[float64](43), shared.Large[float64](47), shared.Large[float64](53))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Large[float64](67), shared.Small[float64](59), shared.Small[float64](61))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Large[float64](73), shared.Small[float64](79), shared.Large[float64](71))
}

func TestMulFloat64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Large[float64](1024), shared.Large[float64](1024), shared.Small[float64](1024))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Large[float64](2048), shared.Large[float64](2048), shared.Large[float64](2048))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Small[float64](4096), shared.Small[float64](4096), shared.Small[float64](4096))
	shared.CheckOperation(t, MulFloat64, fallback.Mul, shared.Small[float64](8192), shared.Small[float64](8192), shared.Large[float64](8192))
}

func TestMulInt32Zero(t *testing.T) {
	shared.CheckOperation(t, MulInt32, fallback.Mul, []int32{}, []int32{}, []int32{})
}

func TestMulInt32Prime(t *testing.T) {
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Small[int32](11), shared.Small[int32](13), shared.Small[int32](17))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Small[int32](29), shared.Small[int32](19), shared.Large[int32](23))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Small[int32](37), shared.Large[int32](41), shared.Small[int32](31))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Small[int32](43), shared.Large[int32](47), shared.Large[int32](53))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Large[int32](67), shared.Small[int32](59), shared.Small[int32](61))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Large[int32](73), shared.Small[int32](79), shared.Large[int32](71))
}

func TestMulInt32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Large[int32](1024), shared.Large[int32](1024), shared.Small[int32](1024))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Large[int32](2048), shared.Large[int32](2048), shared.Large[int32](2048))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Small[int32](4096), shared.Small[int32](4096), shared.Small[int32](4096))
	shared.CheckOperation(t, MulInt32, fallback.Mul, shared.Small[int32](8192), shared.Small[int32](8192), shared.Large[int32](8192))
}

func TestMulInt64Zero(t *testing.T) {
	shared.CheckOperation(t, MulInt64, fallback.Mul, []int64{}, []int64{}, []int64{})
}

func TestMulInt64Prime(t *testing.T) {
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Small[int64](11), shared.Small[int64](13), shared.Small[int64](17))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Small[int64](29), shared.Small[int64](19), shared.Large[int64](23))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Small[int64](37), shared.Large[int64](41), shared.Small[int64](31))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Small[int64](43), shared.Large[int64](47), shared.Large[int64](53))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Large[int64](67), shared.Small[int64](59), shared.Small[int64](61))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Large[int64](73), shared.Small[int64](79), shared.Large[int64](71))
}

func TestMulInt64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Large[int64](1024), shared.Large[int64](1024), shared.Small[int64](1024))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Large[int64](2048), shared.Large[int64](2048), shared.Large[int64](2048))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Small[int64](4096), shared.Small[int64](4096), shared.Small[int64](4096))
	shared.CheckOperation(t, MulInt64, fallback.Mul, shared.Small[int64](8192), shared.Small[int64](8192), shared.Large[int64](8192))
}

func TestOrInt32Zero(t *testing.T) {
	shared.CheckOperation(t, OrInt32, fallback.Sub, []int32{}, []int32{}, []int32{})
}

func TestOrInt32Prime(t *testing.T) {
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Small[int32](11), shared.Small[int32](13), shared.Small[int32](17))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Small[int32](29), shared.Small[int32](19), shared.Large[int32](23))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Small[int32](37), shared.Large[int32](41), shared.Small[int32](31))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Small[int32](43), shared.Large[int32](47), shared.Large[int32](53))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Large[int32](67), shared.Small[int32](59), shared.Small[int32](61))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Large[int32](73), shared.Small[int32](79), shared.Large[int32](71))
}

func TestOrInt32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Large[int32](1024), shared.Large[int32](1024), shared.Small[int32](1024))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Large[int32](2048), shared.Large[int32](2048), shared.Large[int32](2048))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Small[int32](4096), shared.Small[int32](4096), shared.Small[int32](4096))
	shared.CheckOperation(t, OrInt32, fallback.Or, shared.Small[int32](8192), shared.Small[int32](8192), shared.Large[int32](8192))
}

func TestOrInt64Zero(t *testing.T) {
	shared.CheckOperation(t, OrInt64, fallback.Or, []int64{}, []int64{}, []int64{})
}

func TestOrInt64Prime(t *testing.T) {
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Small[int64](11), shared.Small[int64](13), shared.Small[int64](17))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Small[int64](29), shared.Small[int64](19), shared.Large[int64](23))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Small[int64](37), shared.Large[int64](41), shared.Small[int64](31))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Small[int64](43), shared.Large[int64](47), shared.Large[int64](53))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Large[int64](67), shared.Small[int64](59), shared.Small[int64](61))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Large[int64](73), shared.Small[int64](79), shared.Large[int64](71))
}

func TestOrInt64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Large[int64](1024), shared.Large[int64](1024), shared.Small[int64](1024))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Large[int64](2048), shared.Large[int64](2048), shared.Large[int64](2048))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Small[int64](4096), shared.Small[int64](4096), shared.Small[int64](4096))
	shared.CheckOperation(t, OrInt64, fallback.Or, shared.Small[int64](8192), shared.Small[int64](8192), shared.Large[int64](8192))
}

func TestSubFloat32Zero(t *testing.T) {
	shared.CheckOperation(t, SubFloat32, fallback.Sub, []float32{}, []float32{}, []float32{})
}

func TestSubFloat32Prime(t *testing.T) {
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Small[float32](11), shared.Small[float32](13), shared.Small[float32](17))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Small[float32](29), shared.Small[float32](19), shared.Large[float32](23))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Small[float32](37), shared.Large[float32](41), shared.Small[float32](31))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Small[float32](43), shared.Large[float32](47), shared.Large[float32](53))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Large[float32](67), shared.Small[float32](59), shared.Small[float32](61))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Large[float32](73), shared.Small[float32](79), shared.Large[float32](71))
}

func TestSubFloat32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Large[float32](1024), shared.Large[float32](1024), shared.Small[float32](1024))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Large[float32](2048), shared.Large[float32](2048), shared.Large[float32](2048))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Small[float32](4096), shared.Small[float32](4096), shared.Small[float32](4096))
	shared.CheckOperation(t, SubFloat32, fallback.Sub, shared.Small[float32](8192), shared.Small[float32](8192), shared.Large[float32](8192))
}

func TestSubFloat64Zero(t *testing.T) {
	shared.CheckOperation(t, SubFloat64, fallback.Sub, []float64{}, []float64{}, []float64{})
}

func TestSubFloat64Prime(t *testing.T) {
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Small[float64](11), shared.Small[float64](13), shared.Small[float64](17))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Small[float64](29), shared.Small[float64](19), shared.Large[float64](23))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Small[float64](37), shared.Large[float64](41), shared.Small[float64](31))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Small[float64](43), shared.Large[float64](47), shared.Large[float64](53))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Large[float64](67), shared.Small[float64](59), shared.Small[float64](61))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Large[float64](73), shared.Small[float64](79), shared.Large[float64](71))
}

func TestSubFloat64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Large[float64](1024), shared.Large[float64](1024), shared.Small[float64](1024))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Large[float64](2048), shared.Large[float64](2048), shared.Large[float64](2048))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Small[float64](4096), shared.Small[float64](4096), shared.Small[float64](4096))
	shared.CheckOperation(t, SubFloat64, fallback.Sub, shared.Small[float64](8192), shared.Small[float64](8192), shared.Large[float64](8192))
}

func TestSubInt32Zero(t *testing.T) {
	shared.CheckOperation(t, SubInt32, fallback.Sub, []int32{}, []int32{}, []int32{})
}

func TestSubInt32Prime(t *testing.T) {
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Small[int32](11), shared.Small[int32](13), shared.Small[int32](17))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Small[int32](29), shared.Small[int32](19), shared.Large[int32](23))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Small[int32](37), shared.Large[int32](41), shared.Small[int32](31))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Small[int32](43), shared.Large[int32](47), shared.Large[int32](53))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Large[int32](67), shared.Small[int32](59), shared.Small[int32](61))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Large[int32](73), shared.Small[int32](79), shared.Large[int32](71))
}

func TestSubInt32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Large[int32](1024), shared.Large[int32](1024), shared.Small[int32](1024))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Large[int32](2048), shared.Large[int32](2048), shared.Large[int32](2048))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Small[int32](4096), shared.Small[int32](4096), shared.Small[int32](4096))
	shared.CheckOperation(t, SubInt32, fallback.Sub, shared.Small[int32](8192), shared.Small[int32](8192), shared.Large[int32](8192))
}

func TestSubInt64Zero(t *testing.T) {
	shared.CheckOperation(t, SubInt64, fallback.Sub, []int64{}, []int64{}, []int64{})
}

func TestSubInt64Prime(t *testing.T) {
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Small[int64](11), shared.Small[int64](13), shared.Small[int64](17))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Small[int64](29), shared.Small[int64](19), shared.Large[int64](23))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Small[int64](37), shared.Large[int64](41), shared.Small[int64](31))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Small[int64](43), shared.Large[int64](47), shared.Large[int64](53))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Large[int64](67), shared.Small[int64](59), shared.Small[int64](61))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Large[int64](73), shared.Small[int64](79), shared.Large[int64](71))
}

func TestSubInt64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Large[int64](1024), shared.Large[int64](1024), shared.Small[int64](1024))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Large[int64](2048), shared.Large[int64](2048), shared.Large[int64](2048))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Small[int64](4096), shared.Small[int64](4096), shared.Small[int64](4096))
	shared.CheckOperation(t, SubInt64, fallback.Sub, shared.Small[int64](8192), shared.Small[int64](8192), shared.Large[int64](8192))
}

func TestXorInt32Zero(t *testing.T) {
	shared.CheckOperation(t, XorInt32, fallback.Sub, []int32{}, []int32{}, []int32{})
}

func TestXorInt32Prime(t *testing.T) {
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Small[int32](11), shared.Small[int32](13), shared.Small[int32](17))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Small[int32](29), shared.Small[int32](19), shared.Large[int32](23))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Small[int32](37), shared.Large[int32](41), shared.Small[int32](31))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Small[int32](43), shared.Large[int32](47), shared.Large[int32](53))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Large[int32](67), shared.Small[int32](59), shared.Small[int32](61))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Large[int32](73), shared.Small[int32](79), shared.Large[int32](71))
}

func TestXorInt32PowerTwo(t *testing.T) {
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Large[int32](1024), shared.Large[int32](1024), shared.Small[int32](1024))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Large[int32](2048), shared.Large[int32](2048), shared.Large[int32](2048))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Small[int32](4096), shared.Small[int32](4096), shared.Small[int32](4096))
	shared.CheckOperation(t, XorInt32, fallback.Xor, shared.Small[int32](8192), shared.Small[int32](8192), shared.Large[int32](8192))
}

func TestXorInt64Zero(t *testing.T) {
	shared.CheckOperation(t, XorInt64, fallback.Xor, []int64{}, []int64{}, []int64{})
}

func TestXorInt64Prime(t *testing.T) {
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Small[int64](11), shared.Small[int64](13), shared.Small[int64](17))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Small[int64](29), shared.Small[int64](19), shared.Large[int64](23))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Small[int64](37), shared.Large[int64](41), shared.Small[int64](31))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Small[int64](43), shared.Large[int64](47), shared.Large[int64](53))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Large[int64](67), shared.Small[int64](59), shared.Small[int64](61))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Large[int64](73), shared.Small[int64](79), shared.Large[int64](71))
}

func TestXorInt64PowerTwo(t *testing.T) {
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Large[int64](1024), shared.Large[int64](1024), shared.Small[int64](1024))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Large[int64](2048), shared.Large[int64](2048), shared.Large[int64](2048))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Small[int64](4096), shared.Small[int64](4096), shared.Small[int64](4096))
	shared.CheckOperation(t, XorInt64, fallback.Xor, shared.Small[int64](8192), shared.Small[int64](8192), shared.Large[int64](8192))
}

var (
	vectorSizes []int = []int{
		50,
		100,
		200,
		300,
		400,
		500,
		600,
		700,
		800,
		900,
		1000,
		2000,
		3000,
		4000,
		5000,
		6000,
		7000,
		8000,
		9000,
		10000,
		20000,
		30000,
		40000,
		50000,
		60000,
		70000,
		80000,
		90000,
		100000,
	}
)

func BenchmarkSimdAddFloat32(b *testing.B) {
	for _, size := range vectorSizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			left := shared.Small[float32](size)
			right := shared.Small[float32](size)
			result := shared.Small[float32](size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				AddFloat32(left, right, result)
			}
		})
	}
}

func BenchmarkFallbackAddFloat32(b *testing.B) {
	for _, size := range vectorSizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			left := shared.Small[float32](size)
			right := shared.Small[float32](size)
			result := shared.Small[float32](size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				fallback.Add(left, right, result)
			}
		})
	}
}

func Test_AddFloat32(t *testing.T) {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := AddFloat32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []float32{4, 16, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_AddFloat64(t *testing.T) {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := AddFloat64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []float64{4, 16, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_AddInt32(t *testing.T) {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := AddInt32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int32{4, 16, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_AddInt64(t *testing.T) {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := AddInt64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int64{4, 16, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_AndInt32(t *testing.T) {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := AndInt32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int32{1, 1, 0, 0, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_AndInt64(t *testing.T) {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := AndInt64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int64{1, 1, 0, 0, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_DivFloat32(t *testing.T) {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := DivFloat32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []float32{0.33333334, 1.2857143, 0.5, 1.3333334, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_DivFloat64(t *testing.T) {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := DivFloat64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []float64{0.3333333333333333, 1.2857142857142858, 0.5, 1.3333333333333333, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_DivInt32(t *testing.T) {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := DivInt32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int32{0, 1, 0, 1, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_DivInt64(t *testing.T) {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := DivInt64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int64{0, 1, 0, 1, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_OrInt32(t *testing.T) {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := OrInt32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int32{3, 15, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_OrInt64(t *testing.T) {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := OrInt64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int64{3, 15, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_SubFloat32(t *testing.T) {
	left := []float32{1, 9, 2, 8}
	right := []float32{3, 7, 4, 6, 5}
	result := []float32{0, 0, 0, 0, 0, 0}
	length := SubFloat32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []float32{-2, 2, -2, 2, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_SubFloat64(t *testing.T) {
	left := []float64{1, 9, 2, 8}
	right := []float64{3, 7, 4, 6, 5}
	result := []float64{0, 0, 0, 0, 0, 0}
	length := SubFloat64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []float64{-2, 2, -2, 2, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_SubInt32(t *testing.T) {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := SubInt32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int32{-2, 2, -2, 2, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_SubInt64(t *testing.T) {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := SubInt64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int64{-2, 2, -2, 2, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_XorInt32(t *testing.T) {
	left := []int32{1, 9, 2, 8}
	right := []int32{3, 7, 4, 6, 5}
	result := []int32{0, 0, 0, 0, 0, 0}
	length := XorInt32(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int32{2, 14, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func Test_XorInt64(t *testing.T) {
	left := []int64{1, 9, 2, 8}
	right := []int64{3, 7, 4, 6, 5}
	result := []int64{0, 0, 0, 0, 0, 0}
	length := XorInt64(left, right, result)
	assert.Equal(t, len(left), length)
	expected := []int64{2, 14, 6, 14, 0, 0}
	assert.Equal(t, expected, result)
}

func ExampleAddFloat32() {
    left := []float32 {1, 9, 2, 8}
    right := []float32 {3, 7, 4, 6, 5}
    result := []float32 {0, 0, 0, 0, 0, 0}
    length := AddFloat32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [4 16 6 14 0 0]
}

func ExampleAddFloat64() {
    left := []float64 {1, 9, 2, 8}
    right := []float64 {3, 7, 4, 6, 5}
    result := []float64 {0, 0, 0, 0, 0, 0}
    length := AddFloat64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [4 16 6 14 0 0]
}

func ExampleAddInt32() {
    left := []int32 {1, 9, 2, 8}
    right := []int32 {3, 7, 4, 6, 5}
    result := []int32 {0, 0, 0, 0, 0, 0}
    length := AddInt32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [4 16 6 14 0 0]
}

func ExampleAddInt64() {
    left := []int64 {1, 9, 2, 8}
    right := []int64 {3, 7, 4, 6, 5}
    result := []int64 {0, 0, 0, 0, 0, 0}
    length := AddInt64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [4 16 6 14 0 0]
}

func ExampleAndInt32() {
    left := []int32 {1, 9, 2, 8}
    right := []int32 {3, 7, 4, 6, 5}
    result := []int32 {0, 0, 0, 0, 0, 0}
    length := AndInt32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [1 1 0 0 0 0]
}

func ExampleAndInt64() {
    left := []int64 {1, 9, 2, 8}
    right := []int64 {3, 7, 4, 6, 5}
    result := []int64 {0, 0, 0, 0, 0, 0}
    length := AndInt64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [1 1 0 0 0 0]
}

func ExampleDivFloat32() {
    left := []float32 {1, 9, 2, 8}
    right := []float32 {3, 7, 4, 6, 5}
    result := []float32 {0, 0, 0, 0, 0, 0}
    length := DivFloat32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [0.33333334 1.2857143 0.5 1.3333334 0 0]
}

func ExampleDivFloat64() {
    left := []float64 {1, 9, 2, 8}
    right := []float64 {3, 7, 4, 6, 5}
    result := []float64 {0, 0, 0, 0, 0, 0}
    length := DivFloat64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [0.3333333333333333 1.2857142857142858 0.5 1.3333333333333333 0 0]
}

func ExampleDivInt32() {
    left := []int32 {1, 9, 2, 8}
    right := []int32 {3, 7, 4, 6, 5}
    result := []int32 {0, 0, 0, 0, 0, 0}
    length := DivInt32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [0 1 0 1 0 0]
}

func ExampleDivInt64() {
    left := []int64 {1, 9, 2, 8}
    right := []int64 {3, 7, 4, 6, 5}
    result := []int64 {0, 0, 0, 0, 0, 0}
    length := DivInt64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [0 1 0 1 0 0]
}

func ExampleOrInt32() {
    left := []int32 {1, 9, 2, 8}
    right := []int32 {3, 7, 4, 6, 5}
    result := []int32 {0, 0, 0, 0, 0, 0}
    length := OrInt32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [3 15 6 14 0 0]
}

func ExampleOrInt64() {
    left := []int64 {1, 9, 2, 8}
    right := []int64 {3, 7, 4, 6, 5}
    result := []int64 {0, 0, 0, 0, 0, 0}
    length := OrInt64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [3 15 6 14 0 0]
}

func ExampleSubFloat32() {
    left := []float32 {1, 9, 2, 8}
    right := []float32 {3, 7, 4, 6, 5}
    result := []float32 {0, 0, 0, 0, 0, 0}
    length := SubFloat32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [-2 2 -2 2 0 0]
}

func ExampleSubFloat64() {
    left := []float64 {1, 9, 2, 8}
    right := []float64 {3, 7, 4, 6, 5}
    result := []float64 {0, 0, 0, 0, 0, 0}
    length := SubFloat64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [-2 2 -2 2 0 0]
}

func ExampleSubInt32() {
    left := []int32 {1, 9, 2, 8}
    right := []int32 {3, 7, 4, 6, 5}
    result := []int32 {0, 0, 0, 0, 0, 0}
    length := SubInt32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [-2 2 -2 2 0 0]
}

func ExampleSubInt64() {
    left := []int64 {1, 9, 2, 8}
    right := []int64 {3, 7, 4, 6, 5}
    result := []int64 {0, 0, 0, 0, 0, 0}
    length := SubInt64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [-2 2 -2 2 0 0]
}

func ExampleXorInt32() {
    left := []int32 {1, 9, 2, 8}
    right := []int32 {3, 7, 4, 6, 5}
    result := []int32 {0, 0, 0, 0, 0, 0}
    length := XorInt32(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [2 14 6 14 0 0]
}

func ExampleXorInt64() {
    left := []int64 {1, 9, 2, 8}
    right := []int64 {3, 7, 4, 6, 5}
    result := []int64 {0, 0, 0, 0, 0, 0}
    length := XorInt64(left, right, result)
    fmt.Print(length, result)
    // Output: 4 [2 14 6 14 0 0]
}
