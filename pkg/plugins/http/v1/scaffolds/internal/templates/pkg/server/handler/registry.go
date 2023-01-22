package handler

import (
	_ "embed"
	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/utils"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"strings"
	"text/template"
)

//go:embed registry.tmpl
var registerTemplate string

var _ machinery.Template = &Registry{}

// Registry scaffolds a file that defines the config package
type Registry struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Registry) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "handler", strings.ToLower(f.Resource.Group), f.Resource.Version, "registry.go")
	}

	f.TemplateBody = registerTemplate

	return nil
}

func (f *Registry) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}
