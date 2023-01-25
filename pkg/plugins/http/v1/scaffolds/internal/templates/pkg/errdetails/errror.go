package errdetails

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &ErrorDetail{}

// Config scaffolds a file that defines the config package
type ErrorDetail struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *ErrorDetail) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "errdetails", "error.go")
	}

	f.TemplateBody = errdetailsTemplate

	return nil
}

var errdetailsTemplate = `{{ .Boilerplate }}

package errdetails

type BizError struct {
	Code     int               ` + "`" + `json:"code,omitempty"` + "`" + `
	Reason   string            ` + "`" + `json:"reason,omitempty"` + "`" + `
	Metadata map[string]string ` + "`" + `json:"metadata,omitempty"` + "`" + `
}
`
