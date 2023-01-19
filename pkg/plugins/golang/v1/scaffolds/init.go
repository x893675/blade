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

package scaffolds

import (
	"fmt"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/hack"
	configTemplate "github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/config"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/errdetails"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/logger"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/server"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/server/filters"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/server/param"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/server/validate"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/signal"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/utils/sets"
	"github.com/x893675/blade/pkg/plugins/golang/v1/scaffolds/internal/templates/pkg/version"

	"github.com/spf13/afero"

	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugins"
)

const (
	// ControllerRuntimeVersion is the kubernetes-sigs/controller-runtime version to be used in the project
	ControllerRuntimeVersion = "v0.13.1"
)

var _ plugins.Scaffolder = &initScaffolder{}

type initScaffolder struct {
	config          config.Config
	boilerplatePath string
	license         string
	owner           string

	// fs is the filesystem that will be used by the scaffolder
	fs machinery.Filesystem
}

// NewInitScaffolder returns a new Scaffolder for project initialization operations
func NewInitScaffolder(config config.Config, license, owner string) plugins.Scaffolder {
	return &initScaffolder{
		config:          config,
		boilerplatePath: hack.DefaultBoilerplatePath,
		license:         license,
		owner:           owner,
	}
}

// InjectFS implements cmdutil.Scaffolder
func (s *initScaffolder) InjectFS(fs machinery.Filesystem) {
	s.fs = fs
}

// Scaffold implements cmdutil.Scaffolder
func (s *initScaffolder) Scaffold() error {
	fmt.Println("Writing scaffold for you to edit...")

	// Initialize the machinery.Scaffold that will write the boilerplate file to disk
	// The boilerplate file needs to be scaffolded as a separate step as it is going to
	// be used by the rest of the files, even those scaffolded in this command call.
	scaffold := machinery.NewScaffold(s.fs,
		machinery.WithConfig(s.config),
	)

	bpFile := &hack.Boilerplate{
		License: s.license,
		Owner:   s.owner,
	}
	bpFile.Path = s.boilerplatePath
	if err := scaffold.Execute(bpFile); err != nil {
		return err
	}

	boilerplate, err := afero.ReadFile(s.fs.FS, s.boilerplatePath)
	if err != nil {
		return err
	}

	// Initialize the machinery.Scaffold that will write the files to disk
	scaffold = machinery.NewScaffold(s.fs,
		machinery.WithConfig(s.config),
		machinery.WithBoilerplate(string(boilerplate)),
	)

	return scaffold.Execute(
		&templates.Main{},
		&templates.GoMod{},
		&templates.GitIgnore{},
		&templates.Makefile{
			BoilerplatePath: s.boilerplatePath,
		},
		&templates.Dockerfile{},
		&templates.DockerIgnore{},
		&templates.Readme{},
		// pkg/version
		&version.Version{},
		// pkg/utils/sets
		&sets.Sets{},
		&sets.String{},
		// pkg/signal
		&signal.Signal{},
		&signal.POSIXSignal{},
		&signal.WinSignal{},
		// pkg/config
		&configTemplate.Config{},
		// pkg/ent
		//&ent.Generate{},
		//&schema.Base{},
		// pkg/models
		//&models.Models{},
		// pkg/errdetails
		&errdetails.ErrorDetail{},
		// pkg/logger
		&logger.EncodeType{},
		&logger.Logger{},
		&logger.Options{},
		// pkg/mtime
		//&mtime.MTime{},
		// pkg/server
		&filters.Logger{},
		&filters.Skip{},
		&param.Common{},
		&validate.Validate{},
		&server.Server{},
	)
}
