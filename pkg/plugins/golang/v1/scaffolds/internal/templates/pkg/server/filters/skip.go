package filters

import (
	_ "embed"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed skip.tmpl
var skipTemplate string

var _ machinery.Template = &Skip{}

// Skip scaffolds a file that defines the config package
type Skip struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Skip) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "filters", "skip.go")
	}

	f.TemplateBody = skipTemplate

	return nil
}
