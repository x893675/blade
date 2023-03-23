package schema

import (
	_ "embed"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/utils"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed resource.tmpl
var resourceTmpl string

var _ machinery.Template = &Resource{}

type Resource struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin
}

func (r *Resource) SetTemplateDefaults() error {
	if r.Path == "" {
		r.Path = filepath.Join("pkg", "ent", "schema", strings.ToLower(r.Resource.Kind)+".go")
	}

	r.TemplateBody = resourceTmpl

	return nil
}

func (r *Resource) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}
