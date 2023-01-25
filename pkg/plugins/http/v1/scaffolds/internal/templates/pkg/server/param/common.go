package param

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed common.tmpl
var commonTemplate string

var _ machinery.Template = &Common{}

// Common scaffolds a file that defines the config package
type Common struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *Common) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "param", "common.go")
	}

	f.TemplateBody = commonTemplate

	return nil
}
