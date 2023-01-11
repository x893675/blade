package validate

import (
	_ "embed"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed validate.tmpl
var validateTemplate string

var _ machinery.Template = &Validate{}

// Validate scaffolds a file that defines the config package
type Validate struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Validate) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "validate", "validate.go")
	}

	f.TemplateBody = validateTemplate

	return nil
}
