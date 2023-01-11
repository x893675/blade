package sets

import (
	_ "embed"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed string.tmpl
var stringTemplate string

var _ machinery.Template = &String{}

// String scaffolds a file that defines the config package
type String struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *String) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "utils", "sets", "string.go")
	}

	f.TemplateBody = stringTemplate

	return nil
}
