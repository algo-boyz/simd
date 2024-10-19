SHELL := /bin/bash
.DEFAULT_GOAL := build

BIN         = $(CURDIR)/bin
BUILD_DIR   = $(CURDIR)/build

GOPATH      = $(HOME)/go
GOBIN       = $(GOPATH)/bin
GO          ?= GOGC=off $(shell which go)
PKGS        = $(or $(PKG),$(shell env $(GO) list ./...))
VERSION     ?= $(shell git describe --tags --always --match=v*)
SHORT_COMMIT ?= $(shell git rev-parse --short HEAD)
ENVIRONMENT ?= local

PATH := $(BIN):$(GOBIN):$(PATH)

LDFLAGS = -w -s -X "github.com/pehringer/simd/init.version=$(VERSION)" -X "github.com/pehringer/simd/init.commit=$(SHORT_COMMIT)"

# Printing
V ?= 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

$(BUILD_DIR):
	@mkdir -p $@

# Tools
$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(@F)…)
	$Q GOBIN=$(BIN) $(GO) install -v $(shell $(GO) list -tags=tools -e -f '{{ join .Imports "\n" }}' tools/tools.go | grep $(@F))

GOLANGCI_LINT = $(BIN)/golangci-lint
TOOLCHAIN = $(GOLANGCI_LINT)

# Targets
.PHONY: lint # Run project linters
lint: | $(GOLANGCI_LINT)
	$(info $(M) running linter…)
	$Q $(GOLANGCI_LINT) run --max-issues-per-linter 10 --timeout 5m

.PHONY: build # Build service
build: $(BUILD_DIR)
	$(info $(M) $(GOOS) $(GOARCH) building executable…)
	$Q CGO_ENABLED=0 GOOS=$(or $(GOOS),linux) GOARCH=$(or $(GOARCH),amd64) $(GO) build -ldflags '$(LDFLAGS)' -a -o $(BUILD_DIR)/simd ./simd.go
	$Q chmod +x $(BUILD_DIR)/simd
	@true

.PHONY: test # Run all tests and generate coverage report
test:
	$(info $(M) running tests locally…)
	$Q $(GO) test $(PKGS) -race -count 1 -timeout 5m -covermode=atomic -coverprofile cover.out fmt

.PHONY: bench # Run benchmark tests
bench: | $(BUILD_DIR)
	$(info $(M) running benchmarks locally…)
	$Q $(GO) test $(PKGS) -benchmem -run=^\$$ -bench ./...

.PHONY: lint # Run linter inside docker container
docker-lint:
	$(info $(M) docker: running linter…)
	@docker build . --target lint

.PHONY: test-docker # Build simd inside docker container
docker-build:
	$(info $(M) docker: ${TARGETOS} ${TARGETARCH} building executable…)
	@docker build . --target bin \
	--output build/ \
	--platform linux/amd64

.PHONY: test-docker # Run all tests inside docker container
docker-test:
	$(info $(M) docker: running tests…)
	docker-compose up test

.PHONY: unit-test-coverage # Generate coverage report inside docker container
docker-coverage:
	$(info $(M) docker: generate test coverage…)
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: gitest # Run GitHub Actions
gitest:
	$(info $(M) running tests in GitHub Actions…)
	$Q $(GO) test $(PKGS) -race -count 1 -timeout 5m -short

.PHONY: fmt # Run gofmt on go source files
fmt:
	$(info $(M) running go fmt…)
	$Q $(GO) fmt $(PKGS)

.PHONY: generate # Run go generate on go source files
generate: | $(TOOLCHAIN)
	$(info $(M) running go generate…)
	$Q $(GO) generate $(PKGS)
	$Q $(MAKE) fmt

.PHONY: clean # Cleanup project root
clean:
	$(info $(M) cleaning…)
	@rm -rf $(BIN)
	@rm -rf $(BUILD_DIR)

.PHONY: help # Display help
help:
	@grep  -E '^.PHONY' $(MAKEFILE_LIST) | awk 'BEGIN {FS = "#|: "}; {printf "\033[36m%-20s\033[0m %s\n",$$2,$$3}'