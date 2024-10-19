// +build linux darwin
// +build amd64

// func SubFloat64(left, right, result []float64) int
TEXT ·SubFloat64(SB), 4, $0
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
    //Sub four float64 values.
    VMOVUPD (SI)(AX*8), Y0
    VMOVUPD (DX)(AX*8), Y1
    VSUBPD  Y1, Y0, Y2
    VMOVUPD Y2, (DI)(AX*8)
    ADDQ    $4, AX
    JMP     multipleDataLoop
singleDataLoop:
    CMPQ    AX, CX
    JGE     returnLength
    //Sub one float64 value.
    MOVSD   (SI)(AX*8), X0
    MOVSD   (DX)(AX*8), X1
    SUBSD   X1, X0
    MOVSD   X0, (DI)(AX*8)
    INCQ    AX
    JMP     singleDataLoop
returnLength:
    MOVQ    CX, int+72(FP)
    RET
