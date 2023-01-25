package handler

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds/internal/templates/pkg/utils"
)

//go:embed handler.tmpl
var handlerTemplate string

var _ machinery.Template = &Handler{}

// Handler scaffolds a file that defines the config package
type Handler struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.ResourceMixin

	HandlerMarker string
}

// SetTemplateDefaults implements file.Template
func (f *Handler) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "handler.go")
	}

	f.TemplateBody = handlerTemplate
	f.HandlerMarker = machinery.NewMarkerFor(f.Path, "handler").String()

	return nil
}

func (f *Handler) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}

var _ machinery.Inserter = &HandlerUpdater{}

// HandlerUpdater updates main.go to run Controllers
type HandlerUpdater struct { //nolint:maligned
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin
}

// GetPath implements file.Builder
func (f *HandlerUpdater) GetPath() string {
	return getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "handler.go")
}

// GetIfExistsAction implements file.Builder
func (f *HandlerUpdater) GetIfExistsAction() machinery.IfExistsAction {
	return machinery.OverwriteFile
}

// GetMarkers implements file.Inserter
func (f *HandlerUpdater) GetMarkers() []machinery.Marker {
	return []machinery.Marker{
		machinery.NewMarkerFor(getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "handler.go"), "handler"),
	}
}

// GetCodeFragments implements file.Inserter
func (f *HandlerUpdater) GetCodeFragments() machinery.CodeFragmentsMap {
	fragments := make(machinery.CodeFragmentsMap, 1)

	// If resource is not being provided we are creating the file, not updating it
	if f.Resource == nil {
		return fragments
	}

	temp := template.New("handler_updater")
	temp.Funcs(f.GetFuncMap())
	if _, err := temp.Parse(apiHandlerCodeFragment); err != nil {
		fmt.Println("Parse template failed", err.Error())
		return fragments
	}
	out := &bytes.Buffer{}
	if err := temp.Execute(out, f); err != nil {
		fmt.Println("Execute template failed", err.Error())
		return fragments
	}
	// Generate API register code fragments
	handlers := []string{
		out.String(),
	}

	// Only store code fragments in the map if the slices are non-empty
	if len(handlers) != 0 {
		fragments[machinery.NewMarkerFor(getPath(strings.ToLower(f.Resource.Group), f.Resource.Version, "handler.go"), "handler")] = handlers
	}

	return fragments
}

func (f *HandlerUpdater) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"toPascal": utils.ToPascal,
	}
}

const (
	apiHandlerCodeFragment = `// Create{{ toPascal .Resource.Kind }} Create {{ toPascal .Resource.Kind }}
//
//			@Summary      Create {{ toPascal .Resource.Kind }}
//			@Description  Create {{ toPascal .Resource.Kind }}
//			@Tags         {{ .Resource.Group }}
//			@Accept       json
//			@Produce      json
//			@Success      200  {object}  param.PageableResponse{} "desc"
//			@Failure      500  {object}  errdetails.BizError
//			@Router       /{{ lower .Resource.Group }}/{{ .Resource.Version }}/{{ lower .Resource.Kind }} [post]
func (h *handle) Create{{ toPascal .Resource.Kind }}(c echo.Context) error {
	// TODO(user): custom logic here
	return c.String(http.StatusOK, c.Path())
}


// List{{ toPascal .Resource.Kind }} List {{ toPascal .Resource.Kind }}
//
//			@Summary      List {{ toPascal .Resource.Kind }}
//			@Description  list {{ toPascal .Resource.Kind }}
//			@Tags         {{ .Resource.Group }}
//			@Accept       json
//			@Produce      json
//	        @Param        page   query  int  false "pagination" minimum(1) default(1)
//	        @Param        limit   query  int  false "the number of object in response body" minimum(10) default(10)
//			@Success      200  {object}  param.PageableResponse{} "desc"
//			@Failure      500  {object}  errdetails.BizError
//			@Router       /{{ lower .Resource.Group }}/{{ .Resource.Version }}/{{ lower .Resource.Kind }} [get]
func (h *handle) List{{ toPascal .Resource.Kind }}(c echo.Context) error {
	// TODO(user): custom logic here
	return c.String(http.StatusOK, c.Path())
}

// Describe{{ toPascal .Resource.Kind }} Describe {{ toPascal .Resource.Kind }}
//
//			@Summary      Describe {{ toPascal .Resource.Kind }}
//			@Description  Describe {{ toPascal .Resource.Kind }}
//			@Tags         {{ .Resource.Group }}
//			@Accept       json
//			@Produce      json
//          @Param        id     path   int  true  "the resource ID"
//			@Success      200  {object}  param.PageableResponse{} "desc"
//			@Failure      500  {object}  errdetails.BizError
//			@Router       /{{ lower .Resource.Group }}/{{ .Resource.Version }}/{{ lower .Resource.Kind }}/{id} [get]
func (h *handle) Describe{{ toPascal .Resource.Kind }}(c echo.Context) error {
	// TODO(user): custom logic here
	return c.String(http.StatusOK, c.Path())
}

// Update{{ toPascal .Resource.Kind }} Update {{ toPascal .Resource.Kind }}
//
//			@Summary      Update {{ toPascal .Resource.Kind }}
//			@Description  Update {{ toPascal .Resource.Kind }}
//			@Tags         {{ .Resource.Group }}
//			@Accept       json
//			@Produce      json
//          @Param        id     path   int  true  "the resource ID"
//			@Success      200  {object}  param.PageableResponse{} "desc"
//			@Failure      500  {object}  errdetails.BizError
//			@Router       /{{ lower .Resource.Group }}/{{ .Resource.Version }}/{{ lower .Resource.Kind }}/{id} [put]
func (h *handle) Update{{ toPascal .Resource.Kind }}(c echo.Context) error {
	// TODO(user): custom logic here
	return c.String(http.StatusOK, c.Path())
}

// Delete{{ toPascal .Resource.Kind }} Delete {{ toPascal .Resource.Kind }}
//
//			@Summary      Delete {{ toPascal .Resource.Kind }}
//			@Description  Delete {{ toPascal .Resource.Kind }}
//			@Tags         {{ .Resource.Group }}
//			@Accept       json
//			@Produce      json
//          @Param        id     path   int  true  "the resource ID"
//			@Success      200  {object}  param.PageableResponse{} "desc"
//			@Failure      500  {object}  errdetails.BizError
//			@Router       /{{ lower .Resource.Group }}/{{ .Resource.Version }}/{{ lower .Resource.Kind }}/{id} [delete]
func (h *handle) Delete{{ toPascal .Resource.Kind }}(c echo.Context) error {
	// TODO(user): custom logic here
	return c.String(http.StatusOK, c.Path())
}
`
)
