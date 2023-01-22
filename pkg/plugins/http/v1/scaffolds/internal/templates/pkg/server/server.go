package server

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"strings"
)

//go:embed server.tmpl
var serverTemplate string

var _ machinery.Template = &Server{}

var defaultFilePath = filepath.Join("pkg", "server", "server.go")

var (
	importMarker = machinery.NewMarkerFor(defaultFilePath, "imports")
	routerMarker = machinery.NewMarkerFor(defaultFilePath, "routers")
)

// Server scaffolds a file that defines the config package
type Server struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ProjectNameMixin

	ImportMarker string
	RouterMarker string
}

// SetTemplateDefaults implements file.Template
func (f *Server) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = defaultFilePath
	}
	f.ImportMarker = importMarker.String()
	f.RouterMarker = routerMarker.String()
	f.TemplateBody = serverTemplate

	return nil
}

var _ machinery.Inserter = &ServerUpdater{}

// ServerUpdater updates main.go to run Controllers
type ServerUpdater struct { //nolint:maligned
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ProjectNameMixin
	machinery.ResourceMixin
}

// GetPath implements file.Builder
func (*ServerUpdater) GetPath() string {
	return defaultFilePath
}

// GetIfExistsAction implements file.Builder
func (*ServerUpdater) GetIfExistsAction() machinery.IfExistsAction {
	return machinery.OverwriteFile
}

// GetMarkers implements file.Inserter
func (f *ServerUpdater) GetMarkers() []machinery.Marker {
	return []machinery.Marker{
		importMarker,
		routerMarker,
	}
}

// GetCodeFragments implements file.Inserter
func (f *ServerUpdater) GetCodeFragments() machinery.CodeFragmentsMap {
	fragments := make(machinery.CodeFragmentsMap, 2)

	// If resource is not being provided we are creating the file, not updating it
	if f.Resource == nil {
		return fragments
	}

	// Generate import code fragments
	imports := []string{
		fmt.Sprintf(routerImportCodeFragment,
			f.Resource.PackageName(), f.Resource.Version, f.Repo, strings.ToLower(f.Resource.Group), f.Resource.Version),
	}

	routers := []string{
		fmt.Sprintf(routerRegisterCodeFragment,
			f.Resource.PackageName(), f.Resource.Version),
	}

	// Only store code fragments in the map if the slices are non-empty
	if len(imports) != 0 {
		fragments[importMarker] = imports
	}

	if len(routers) != 0 {
		fragments[routerMarker] = routers
	}

	return fragments
}

const (
	routerImportCodeFragment = `%s%s "%s/pkg/server/handler/%s/%s"
`
	routerRegisterCodeFragment = `%s%s.RegisterRouter(s.e)
`
)
