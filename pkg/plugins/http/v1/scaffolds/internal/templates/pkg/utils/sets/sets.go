package sets

import (
	_ "embed"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed sets.tmpl
var setsTemplate string

var _ machinery.Template = &Sets{}

// Sets scaffolds a file that defines the config package
type Sets struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *Sets) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "utils", "sets", "sets.go")
	}

	f.TemplateBody = setsTemplate

	return nil
}
