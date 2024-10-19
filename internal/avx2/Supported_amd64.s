// +build linux darwin
// +build amd64

// func Supported() bool
TEXT ·Supported(SB), 4, $0
    //Check XSAVE supported.
    MOVQ    $1, AX
    CPUID
    TESTQ   $(1<<26), CX
    JZ      avx2False
    //Check AVX2 supported.
    MOVQ    $7, AX
    MOVQ    $0, CX
    CPUID
    TESTQ   $(1<<5), BX
    JZ      avx2False
    //Check XMM and YMM enabled.
    MOVQ    $0, CX
    XGETBV
    TESTQ   $(1<<1), AX
    JZ      avx2False
    TESTQ   $(1<<2), AX
    JZ      avx2False
    //avx2True:
    MOVB    $1, bool+0(FP)
    RET
avx2False:
    MOVB    $0, bool+0(FP)
    RET
