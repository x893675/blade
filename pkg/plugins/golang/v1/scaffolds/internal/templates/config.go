package templates

import (
	_ "embed"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed config.tmpl
var configTemplate string

var _ machinery.Template = &Config{}

// Config scaffolds a file that defines the config package
type Config struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *Config) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "config.yaml"
	}

	f.TemplateBody = configTemplate

	return nil
}
