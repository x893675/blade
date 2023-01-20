package signal

import (
	_ "embed"
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

//go:embed signal_windows.tmpl
var signalWINTemplate string

// WinSignal scaffolds a file that defines the config package
type WinSignal struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *WinSignal) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "signal", "signal_windows.go")
	}

	f.TemplateBody = signalWINTemplate

	return nil
}
