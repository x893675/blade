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
	"fmt"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

const defaultMainPath = "main.go"

var _ machinery.Template = &Main{}

// Main scaffolds a file that defines the controller manager entry point
type Main struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ProjectNameMixin
}

// SetTemplateDefaults implements file.Template
func (f *Main) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(defaultMainPath)
	}

	f.TemplateBody = fmt.Sprintf(mainTemplate,
		machinery.NewMarkerFor(f.Path, importMarker),
	)

	return nil
}

var _ machinery.Inserter = &MainUpdater{}

// MainUpdater updates main.go to run Controllers
type MainUpdater struct { //nolint:maligned
	machinery.RepositoryMixin
	machinery.ProjectNameMixin
}

// GetPath implements file.Builder
func (*MainUpdater) GetPath() string {
	return defaultMainPath
}

// GetIfExistsAction implements file.Builder
func (*MainUpdater) GetIfExistsAction() machinery.IfExistsAction {
	return machinery.OverwriteFile
}

const (
	importMarker = "imports"
)

// GetMarkers implements file.Inserter
func (f *MainUpdater) GetMarkers() []machinery.Marker {
	return []machinery.Marker{
		machinery.NewMarkerFor(defaultMainPath, importMarker),
	}
}

// GetCodeFragments implements file.Inserter
func (f *MainUpdater) GetCodeFragments() machinery.CodeFragmentsMap {
	fragments := make(machinery.CodeFragmentsMap, 1)

	// Generate import code fragments
	imports := make([]string, 0)

	// Only store code fragments in the map if the slices are non-empty
	if len(imports) != 0 {
		fragments[machinery.NewMarkerFor(defaultMainPath, importMarker)] = imports
	}

	return fragments
}

var mainTemplate = `{{ .Boilerplate }}

package main

import (
	"flag"
	"fmt"
	"{{ .Repo }}/pkg/config"
	"{{ .Repo }}/pkg/logger"
	"{{ .Repo }}/pkg/server"
	"{{ .Repo }}/pkg/signal"
	"{{ .Repo }}/pkg/version"
	"math/rand"
	"os"
	"time"

	_ "{{ .Repo }}/docs"
	_ "{{ .Repo }}/pkg/ent/runtime"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	%s
)

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "{{ .ProjectName }}: ")
		// TODO(user): add custom project long description
		_, _ = fmt.Fprintf(os.Stderr, "Custom Project Long Description.\n\n")
		_, _ = fmt.Fprintln(os.Stderr, "VERSION:  ", version.Version)
		flag.PrintDefaults()
	}
}

var (
	configPath = flag.String("config", "config.yaml", "The server config file path")
)

// @title {{ .ProjectName }} API
// @version v0.0.1
// @description This is a {{ .ProjectName }} server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host locahost:8080
// @BasePath /

func main() {
	flag.Parse()
	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	logger.ApplyZapLoggerWithOptions(cfg.LogOptions)

	ctx := signal.SetupSignalContext()

	s := server.NewService(cfg)

	if err := s.PrepareRun(ctx); err != nil {
		panic(err)
	}

	if err := s.Run(ctx); err != nil {
		panic(err)
	}
}
`
