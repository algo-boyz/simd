// +build linux darwin
// +build arm64

#include "textflag.h"

TEXT ·AddFloat32(SB), $0-72
    MOVD    left_len+8(FP), R0    // Load left length
    MOVD    right_len+32(FP), R1  // Load right length
    MOVD    result_len+56(FP), R2 // Load result length

    // Find minimum length
    CMP     R0, R1
    BHI     use_r1               // If R0 > R1, use R1
    MOVD    R0, R1               // else keep R0 in R1
use_r1:
    CMP     R1, R2
    BLS     setup                // If R1 <= R2, continue
    MOVD    R2, R1               // else use R2

setup:
    CBZ     R1, done            // If length is 0, return

    MOVD    left_base+0(FP), R3   // Load left base
    MOVD    right_base+24(FP), R4 // Load right base
    MOVD    result_base+48(FP), R5 // Load result base
    MOVD    ZR, R6                // Initialize index

loop:
    FMOVS   (R3)(R6), F0         // Load from left
    FMOVS   (R4)(R6), F1         // Load from right
    FADDS   F0, F1, F2           // Add floats
    FMOVS   F2, (R5)(R6)         // Store result

    ADD     $4, R6                // Increment index by 4 (Plan9 syntax)
    MOVD    R1, R7                // Copy length to R7
    LSL     $2, R7                // Shift left by 2 (Plan9 syntax)
    CMP     R6, R7                // Compare index with (length * 4)
    BLT     loop                 // Continue if index < (length * 4)

done:
    MOVD    R1, ret+64(FP)       // Return minimum length
    RET
