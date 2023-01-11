package config

import (
	"path/filepath"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &Config{}

// Config scaffolds a file that defines the config package
type Config struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Config) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "config", "config.go")
	}

	f.TemplateBody = configTemplate

	return nil
}

var configTemplate = `{{ .Boilerplate }}

package config

import (
	"{{ .Repo }}/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Debug                   bool              ` + "`" + `json:"debug" yaml:"debug"` + "`" + `
	GenericServerRunOptions *ServerRunOptions ` + "`" + `json:"generic" yaml:"generic"` + "`" + `
	LogOptions              *logger.Options   ` + "`" + `json:"log,omitempty" yaml:"log,omitempty"` + "`" + `
	Database                *Database         ` + "`" + `json:"db" yaml:"db"` + "`" + `
}

type ServerRunOptions struct {
	BindAddress   string ` + "`" + `json:"bindAddress" yaml:"bindAddress"` + "`" + `
	Port          int    ` + "`" + `json:"port" yaml:"port"` + "`" + `
	TLSCertFile   string ` + "`" + `json:"tlsCertFile" yaml:"tlsCertFile"` + "`" + `
	TLSPrivateKey string ` + "`" + `json:"tlsPrivateKey" yaml:"tlsPrivateKey"` + "`" + `
}

type Database struct {
	Driver string ` + "`" + `json:"driver" yaml:"driver"` + "`" + `
	Dsn    string ` + "`" + `json:"dsn" yaml:"dsn"` + "`" + `
}

func Load(path string) (*Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(f, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}`
