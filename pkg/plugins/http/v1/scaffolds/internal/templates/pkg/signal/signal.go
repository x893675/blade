package signal

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed signal.tmpl
var signalTemplate string

var _ machinery.Template = &Signal{}

// Signal scaffolds a file that defines the config package
type Signal struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *Signal) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "signal", "signal.go")
	}

	f.TemplateBody = signalTemplate

	return nil
}
