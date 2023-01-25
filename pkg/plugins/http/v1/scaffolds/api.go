package scaffolds

import (
	"fmt"

	"github.com/spf13/afero"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/model/resource"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugins"

	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/hack"
	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/server"
	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/server/handler"
)

var _ plugins.Scaffolder = &apiScaffolder{}

// apiScaffolder contains configuration for generating scaffolding for Go type
// representing the API and controller that implements the behavior for the API.
type apiScaffolder struct {
	config   config.Config
	resource resource.Resource

	// fs is the filesystem that will be used by the scaffolder
	fs machinery.Filesystem

	// force indicates whether to scaffold controller files even if it exists or not
	force bool
}

// NewAPIScaffolder returns a new Scaffolder for API/controller creation operations
func NewAPIScaffolder(config config.Config, res resource.Resource, force bool) plugins.Scaffolder {
	return &apiScaffolder{
		config:   config,
		resource: res,
		force:    force,
	}
}

// InjectFS implements cmdutil.Scaffolder
func (s *apiScaffolder) InjectFS(fs machinery.Filesystem) {
	s.fs = fs
}

// Scaffold implements cmdutil.Scaffolder
func (s *apiScaffolder) Scaffold() error {
	if s.config.HasResource(s.resource.GVK) && !s.force {
		fmt.Println("Resource scaffold exist, skip now")
		return nil
	}

	var doUpdate bool
	if s.config.HasGroupVersion(s.resource.Group, s.resource.Version) {
		doUpdate = true
	}

	fmt.Println("Writing scaffold for you to edit...")

	// Load the boilerplate
	boilerplate, err := afero.ReadFile(s.fs.FS, hack.DefaultBoilerplatePath)
	if err != nil {
		return fmt.Errorf("error scaffolding API/controller: unable to load boilerplate: %w", err)
	}

	// Initialize the machinery.Scaffold that will write the files to disk
	scaffold := machinery.NewScaffold(s.fs,
		machinery.WithConfig(s.config),
		machinery.WithBoilerplate(string(boilerplate)),
		machinery.WithResource(&s.resource),
	)

	if err := s.config.UpdateResource(s.resource); err != nil {
		return fmt.Errorf("error updating resource: %w", err)
	}

	if doUpdate {
		if err := scaffold.Execute(
			&handler.RegistryUpdater{},
			&handler.HandlerUpdater{},
		); err != nil {
			return fmt.Errorf("error scaffolding APIs: %v", err)
		}
		return nil
	}
	if err := scaffold.Execute(
		&handler.Base{},
		&handler.Registry{},
		&handler.Handler{},
	); err != nil {
		return fmt.Errorf("error scaffolding APIs: %v", err)
	}

	if err := scaffold.Execute(&server.ServerUpdater{}); err != nil {
		return fmt.Errorf("error updating server.go: %v", err)
	}
	return nil
}
