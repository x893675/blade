package mtime

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed mtime.tmpl
var timeTemplate string

var _ machinery.Template = &MTime{}

// MTime scaffolds a file that defines the config package
type MTime struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *MTime) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "mtime", "time.go")
	}

	f.TemplateBody = timeTemplate

	return nil
}
