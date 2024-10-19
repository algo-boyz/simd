// +build linux darwin
// +build amd64

// func Supported() bool
TEXT ·Supported(SB), NOSPLIT, $0-1
    // Read ID_AA64ISAR0_EL1 system register
    // MRS instruction contains information about supported CPU features on ARM64
    MRS     ID_AA64ISAR0_EL1, R0
    // Check if AES feature is supported (bits 7:4 of MRS register)
    // if bits are non-zero, AES is supported.
    AND     $0xF0, R0, R1
    CMP     $0, R1
    BEQ     notSupported
    // Feature is supported
    MOVD    $1, R0
    MOVB    R0, ret+0(FP)
    RET

notSupported:
    MOVD    $0, R0
    MOVB    R0, ret+0(FP)
    RET
    