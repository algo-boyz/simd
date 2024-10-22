//go:build arm64
// +build arm64

package simd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pehringer/simd/internal/neon"
	"golang.org/x/sys/cpu"
)

var (
	welcome = fmt.Sprintf("🏴‍☠ simd version: %s, commit: %s, arch: %s\n", Version(), Commit(), runtime.GOARCH)
	useNEON = cpu.ARM64.HasASIMD
	useSVE  = cpu.ARM64.HasSVE
	// todo https://github.com/minio/sha256-simd/blob/master/cpuid_other.go
)

func init() {
	fmt.Fprint(os.Stdout, welcome)

	if useNEON || useSVE {
		addFloat32 = neon.AddFloat32
		// mulFloat64 = neon.MulFloat64
		return
	}
}
