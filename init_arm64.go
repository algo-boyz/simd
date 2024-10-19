//go:build arm64
// +build arm64

package simd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pehringer/simd/internal/neon"
)

var welcome = fmt.Sprintf("🏴‍☠ simd version: %s, commit: %s, arch: %s\n", Version(), Commit(), runtime.GOARCH)

func init() {
	fmt.Fprint(os.Stdout, welcome)

	if neon.Supported() {
		// addFloat32 = neon.AddFloat32
		return
	}
}
