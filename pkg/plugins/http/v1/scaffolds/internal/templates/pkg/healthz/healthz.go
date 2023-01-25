package healthz

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed healthz.tmpl
var healthCheckTemplate string

var _ machinery.Template = &Healthz{}

var defaultFilePath = filepath.Join("pkg", "healthz", "healthz.go")

// Healthz scaffolds a file that defines the config package
type Healthz struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Healthz) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = defaultFilePath
	}

	f.TemplateBody = healthCheckTemplate

	return nil
}
