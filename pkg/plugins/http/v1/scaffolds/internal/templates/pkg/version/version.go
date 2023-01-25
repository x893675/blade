package version

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed version.tmpl
var versionTemplate string

var _ machinery.Template = &Version{}

// Version scaffolds a file that defines the config package
type Version struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *Version) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "version", "version.go")
	}

	f.TemplateBody = versionTemplate

	return nil
}
