// nolint
package shared

type (
	Floating                        interface{ float32 | float64 }
	Integer                         interface{ int32 | int64 }
	Operation[T Floating | Integer] func(left, right, result []T) int
)
