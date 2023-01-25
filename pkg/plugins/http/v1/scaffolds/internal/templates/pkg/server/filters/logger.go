package filters

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed logger.tmpl
var loggerTemplate string

var _ machinery.Template = &Logger{}

// Logger scaffolds a file that defines the config package
type Logger struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Logger) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "filters", "server.go")
	}

	f.TemplateBody = loggerTemplate

	return nil
}
