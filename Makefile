# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

GOHOSTOS=$(shell go env GOHOSTOS)
GOHOSTARCH=$(shell go env GOHOSTARCH)

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
GORELEASER ?= $(LOCALBIN)/goreleaser
GOIMPORTS ?= $(LOCALBIN)/goimports
GOLINT ?= $(LOCALBIN)/golangci-lint
## Tool Versions
GORELEASER_VERSION ?= v1.14.1
GOIMPORTS_VERSION ?= latest
GOLINT_VERSION ?= v1.50.1

.PHONY: goreleaser
goreleaser: $(LOCALBIN) ## Download gorelaser locally if necessary. If wrong version is installed, it will be overwritten.
	test -s $(LOCALBIN)/goreleaser && $(LOCALBIN)/goreleaser --version | grep -q $(GORELEASER_VERSION) || \
	GOBIN=$(LOCALBIN) go install github.com/goreleaser/goreleaser@$(GORELEASER_VERSION)
	$(GORELEASER) check

.PHONY: goimports
goimports: $(LOCALBIN) ## Download goimports locally if necessary.
	test -s $(LOCALBIN)/goimports || GOBIN=$(LOCALBIN) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)
	@$(GOIMPORTS) -w -local github.com/x893675/blade $(shell find . -type f -name '*.go' -not -path "./staging/*")

.PHONY: golint
golint: $(LOCALBIN) ## Download gorelaser locally if necessary. If wrong version is installed, it will be overwritten.
	test -s $(LOCALBIN)/golangci-lint && $(LOCALBIN)/golangci-lint --version | grep -q $(GOLINT_VERSION) || \
	GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINT_VERSION)
	$(GOLINT) run --timeout 10m

##@ Development

.PHONY: fmt
fmt: goimports## Run go fmt, goimports against code.
	@go fmt $(shell go list ./... | grep -v /staging/)

##@ Build

.PHONY: build
build: fmt golint goreleaser ## Build binary with local GOOS and GOARCH.
	$(GORELEASER) build -f .build.yaml --snapshot --rm-dist --id local-$(GOHOSTOS)-$(GOHOSTARCH)

.PHONY: build-release
build-release: fmt golint goreleaser ## Build binary with release mode.
	$(GORELEASER) build --snapshot --rm-dist --id release

