package ent

import (
	_ "embed"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed generate.tmpl
var generateTemplate string

var _ machinery.Template = &Generate{}

// Generate scaffolds a file that defines the config package
type Generate struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *Generate) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "ent", "generate.go")
	}

	f.TemplateBody = generateTemplate

	return nil
}
