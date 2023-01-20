package logger

import (
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &Options{}

// Options scaffolds a file that defines the config package
type Options struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *Options) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "logger", "options.go")
	}

	f.TemplateBody = optionsTemplate

	return nil
}

var optionsTemplate = `{{ .Boilerplate }}

package logger

type Options struct {
	LogFile          string ` + "`" + `json:"logFile" yaml:"logFile"` + "`" + `
	LogFileMaxSizeMB int    ` + "`" + `json:"logFileMaxSizeMB" yaml:"logFileMaxSizeMB"` + "`" + `
	ToStderr         bool   ` + "`" + `json:"toStderr" yaml:"toStderr"` + "`" + `
	Level            string ` + "`" + `json:"level" yaml:"level"` + "`" + `
	EncodeType       string ` + "`" + `json:"encodeType" yaml:"encodeType"` + "`" + `
	MaxBackups       int    ` + "`" + `json:"maxBackups" yaml:"maxBackups"` + "`" + `
	MaxAge           int    ` + "`" + `json:"maxAge" yaml:"maxAge"` + "`" + `
	Compress         bool   ` + "`" + `json:"compress" yaml:"compress"` + "`" + `
	UseLocalTimeBack bool   ` + "`" + `json:"useLocalTime" yaml:"useLocalTime"` + "`" + `
}

func NewLogOptions() *Options {
	return &Options{
		ToStderr:         true,
		LogFile:          "",
		LogFileMaxSizeMB: 100,
		Level:            "info",
		EncodeType:       "console",
		MaxAge:           30,
		MaxBackups:       5,
		Compress:         false,
		UseLocalTimeBack: true,
	}
}
`
