// +build linux darwin
// +build amd64

// func AndInt64(left, right, result []int64) int
TEXT ·AndInt64(SB), 4, $0
    //Load slices lengths.
    MOVQ    leftLen+8(FP), AX
    MOVQ    rightLen+32(FP), BX
    MOVQ    resultLen+56(FP), CX
    //Get minimum length.
    CMPQ    AX, CX
    CMOVQLT AX, CX
    CMPQ    BX, CX
    CMOVQLT BX, CX
    //Load slices data pointers.
    MOVQ    leftData+0(FP), SI
    MOVQ    rightData+24(FP), DX
    MOVQ    resultData+48(FP), DI
    //Initialize loop index.
    MOVQ    $0, AX
multipleDataLoop:
    MOVQ    CX, BX
    SUBQ    AX, BX
    CMPQ    BX, $4
    JL      singleDataLoop
    //And four int64 values.
    VMOVDQU (SI)(AX*8), Y0
    VMOVDQU (DX)(AX*8), Y1
    VPAND   Y1, Y0, Y2
    VMOVDQU Y2, (DI)(AX*8)
    ADDQ    $4, AX
    JMP     multipleDataLoop
singleDataLoop:
    CMPQ    AX, CX
    JGE     returnLength
    //And one int64 value.
    MOVQ    (SI)(AX*8), R8
    MOVQ    (DX)(AX*8), R9
    ANDQ    R9, R8
    MOVQ    R8, (DI)(AX*8)
    INCQ    AX
    JMP     singleDataLoop
returnLength:
    MOVQ    CX, int+72(FP)
    RET
