package handler

import (
	_ "embed"
	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/utils"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"strings"
	"text/template"
)

//go:embed handler.tmpl
var handlerTemplate string

var _ machinery.Template = &Handler{}

// Handler scaffolds a file that defines the config package
type Handler struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.ResourceMixin
}

// SetTemplateDefaults implements file.Template
func (f *Handler) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "server", "handler", strings.ToLower(f.Resource.Group), f.Resource.Version, "handler.go")
	}

	f.TemplateBody = handlerTemplate

	return nil
}

func (f *Handler) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}
