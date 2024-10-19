package simd

// version of the compiled binaries - variable is
// populated while building through LD flags.
var version string

// Version returns the version of the service. <unknown> is returned if version is not set.
func Version() string {
	if version == "" {
		return "<unknown>"
	}
	return version
}

// commit hash of the compiled binaries.
// set at runtime from build info
var commit string

// Commit hash of the serivice. <unknown> is returned if commit hash is not set.
func Commit() string {
	if commit == "" {
		return "<unknown>"
	}
	return commit
}
