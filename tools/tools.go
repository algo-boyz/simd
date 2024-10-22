//go:build tools
// +build tools

package main

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/minio/asm2plan9s"
	_ "github.com/stretchr/testify"
	_ "github.com/testcontainers/testcontainers-go"
	_ "github.com/testcontainers/testcontainers-go/modules/ollama"
	_ "github.com/tmc/langchaingo/llms"
)
