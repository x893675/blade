package server

import (
	_ "embed"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed server.tmpl
var serverTemplate string

var _ machinery.Template = &Server{}

// Server scaffolds a file that defines the config package
type Server struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ProjectNameMixin
}

// SetTemplateDefaults implements file.Template
func (f *Server) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "server.go")
	}

	f.TemplateBody = serverTemplate

	return nil
}
