package signal

import (
	_ "embed"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed signal_posix.tmpl
var signalPOSIXTemplate string

// POSIXSignal scaffolds a file that defines the config package
type POSIXSignal struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *POSIXSignal) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "signal", "signal_posix.go")
	}

	f.TemplateBody = signalPOSIXTemplate

	return nil
}
