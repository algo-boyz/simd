// +build linux darwin
// +build arm64

// func Supported() bool
TEXT ·Supported(SB), 4, $0
//     // check arm64 support
//     MOVQ    $1, AX
//     CPUID
//     TESTQ   $(1<<25), DX
//     JZ      armFalse
//     // armTrue:
//     MOVB    $1, bool+0(FP) // return 1 if supported
//     RET  // ret+0(FP) is go's equivalent of bool
// armFalse:
    MOVB    $0, bool+0(FP) // return 0 if not supported
    RET
