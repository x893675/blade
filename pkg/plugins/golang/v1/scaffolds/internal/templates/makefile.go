/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package templates

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &Makefile{}

// Makefile scaffolds a file that defines project management CLI commands
type Makefile struct {
	machinery.TemplateMixin
	machinery.ComponentConfigMixin
	machinery.ProjectNameMixin

	// BoilerplatePath is the path to the boilerplate file
	BoilerplatePath string
}

// SetTemplateDefaults implements file.Template
func (f *Makefile) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "Makefile"
	}

	f.TemplateBody = makefileTemplate

	f.IfExistsAction = machinery.Error

	return nil
}

//nolint:lll
const makefileTemplate = `
# default docker image url
IMG ?= {{ .ProjectName }}
# container build tools, default is docker, you can overwrite to nerdctl or other
BUILD_TOOL ?= docker
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

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
SWAG ?= $(LOCALBIN)/swag

## Tool Versions
SWAG_VERSION ?= v1.8.9

.PHONY: swagger
swagger: $(SWAG) ## Download swag locally if necessary. If wrong version is installed, it will be overwritten.
$(SWAG): $(LOCALBIN)
	test -s $(LOCALBIN)/swag && $(LOCALBIN)/swag --version | grep -q $(SWAG_VERSION) || \
	GOBIN=$(LOCALBIN) go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)

##@ Development

.PHONY: generate
generate: swagger ## Generate code containing Swagger, Ent etc.
	$(SWAG) init --parseDependency
	go generate ./pkg/...

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

##@ Build

.PHONY: build
build: generate fmt vet ## Build binary.
	CGO_ENABLED=0 go build -ldflags '-s -w' -o bin/{{ .ProjectName }} main.go

.PHONY: run
run: generate fmt vet ## Run a server from your host.
	go run ./main.go

.PHONY: build-image
docker-build: ## Build docker image with the {{ .ProjectName }}.
	${BUILD_TOOL} build -t ${IMG} .

`
