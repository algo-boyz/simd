// +build linux darwin
// +build amd64

// func MulFloat32(left, right, result []float32) int
TEXT ·MulFloat32(SB), 4, $0
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
    //Mul four float32 values.
    MOVUPS  (SI)(AX*4), X0
    MOVUPS  (DX)(AX*4), X1
    MULPS   X1, X0
    MOVUPS  X0, (DI)(AX*4)
    ADDQ    $4, AX
    JMP     multipleDataLoop
singleDataLoop:
    CMPQ    AX, CX
    JGE     returnLength
    //Mul one float32 value.
    MOVSS   (SI)(AX*4), X0
    MOVSS   (DX)(AX*4), X1
    MULSS   X1, X0
    MOVSS   X0, (DI)(AX*4)
    INCQ    AX
    JMP     singleDataLoop
returnLength:
    MOVQ    CX, int+72(FP)
    RET
