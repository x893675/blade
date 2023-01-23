package handler

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/utils"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"strings"
	"text/template"
)

//go:embed registry.tmpl
var registerTemplate string

func getPath(group, version, filename string) string {
	return filepath.Join("pkg", "server", "handler", group, version, filename)
}

var _ machinery.Template = &Registry{}

// Registry scaffolds a file that defines the config package
type Registry struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.ResourceMixin
	machinery.RepositoryMixin

	RegistryMarker string
}

// SetTemplateDefaults implements file.Template
func (f *Registry) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "registry.go")
	}

	f.TemplateBody = registerTemplate
	f.RegistryMarker = machinery.NewMarkerFor(f.Path, "registry").String()

	return nil
}

func (f *Registry) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}

var _ machinery.Inserter = &RegistryUpdater{}

// RegistryUpdater updates main.go to run Controllers
type RegistryUpdater struct { //nolint:maligned
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin
}

// GetPath implements file.Builder
func (f *RegistryUpdater) GetPath() string {
	return getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "registry.go")
}

// GetIfExistsAction implements file.Builder
func (f *RegistryUpdater) GetIfExistsAction() machinery.IfExistsAction {
	return machinery.OverwriteFile
}

// GetMarkers implements file.Inserter
func (f *RegistryUpdater) GetMarkers() []machinery.Marker {
	return []machinery.Marker{
		machinery.NewMarkerFor(getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "registry.go"), "registry"),
	}
}

// GetCodeFragments implements file.Inserter
func (f *RegistryUpdater) GetCodeFragments() machinery.CodeFragmentsMap {
	fragments := make(machinery.CodeFragmentsMap, 1)

	// If resource is not being provided we are creating the file, not updating it
	if f.Resource == nil {
		return fragments
	}

	temp := template.New("registry_updater")
	temp.Funcs(f.GetFuncMap())
	if _, err := temp.Parse(apiRegisterCodeFragment); err != nil {
		fmt.Println("Parse template failed", err.Error())
		return fragments
	}
	out := &bytes.Buffer{}
	if err := temp.Execute(out, f); err != nil {
		fmt.Println("Execute template failed", err.Error())
		return fragments
	}
	// Generate API register code fragments
	registers := []string{
		out.String(),
	}

	// Only store code fragments in the map if the slices are non-empty
	if len(registers) != 0 {
		fragments[machinery.NewMarkerFor(getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "registry.go"), "registry")] = registers
	}

	return fragments
}

func (f *RegistryUpdater) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}

const (
	apiRegisterCodeFragment = `g.POST("/{{ lower .Resource.Kind }}", h.Create{{ toPascal .Resource.Kind }})
g.GET("/{{ lower .Resource.Kind }}", h.List{{ toPascal .Resource.Kind }})
g.GET("/{{ lower .Resource.Kind }}/:id", h.Describe{{ toPascal .Resource.Kind }})
g.PUT("/{{ lower .Resource.Kind }}/:id", h.Update{{ toPascal .Resource.Kind }})
g.DELETE("/{{ lower .Resource.Kind }}/:id", h.Delete{{ toPascal .Resource.Kind }})
`
)
