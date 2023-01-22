package handler

import (
	_ "embed"
	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/utils"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"strings"
	"text/template"
)

//go:embed base.tmpl
var baseTemplate string

var _ machinery.Template = &Base{}

// Base scaffolds a file that defines the config package
type Base struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.ResourceMixin
}

// SetTemplateDefaults implements file.Template
func (f *Base) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "handler", strings.ToLower(f.Resource.Group), "docs.go")
	}

	f.TemplateBody = baseTemplate

	return nil
}

func (f *Base) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}
