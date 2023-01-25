package schema

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed base.tmpl
var baseTemplate string

var _ machinery.Template = &Base{}

// Base scaffolds a file that defines the config package
type Base struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Base) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "ent", "schema", "base.go")
	}

	f.TemplateBody = baseTemplate

	return nil
}
