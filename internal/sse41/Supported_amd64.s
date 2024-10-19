// +build linux darwin
// +build amd64

// func Supported() bool
TEXT ·Supported(SB), 4, $0
    //Check SSE4.1 supported.
    MOVQ    $1, AX
    CPUID
    TESTQ   $(1<<19), CX
    JZ      sse41False
    //sse41True:
    MOVB    $1, bool+0(FP)
    RET
sse41False:
    MOVB    $0, bool+0(FP)
    RET
