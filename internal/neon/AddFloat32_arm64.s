// +build linux darwin
// +build arm64

// func AddFloat32(left, right, result []float32) int
// ARM64 uses NEON SIMD instructions. 
// We use LDP (Load Pair) and STP (Store Pair) to load and store 4 float32 values at once, and 
// FADD instruction for vector addition using registers R0-R7 / V0-V3
// TEXT ·AddFloat32(SB), $0-72