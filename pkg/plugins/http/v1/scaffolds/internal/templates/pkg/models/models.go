package models

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &Models{}

// Models scaffolds a file that defines the config package
type Models struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
}

// SetTemplateDefaults implements file.Template
func (f *Models) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("pkg", "models", "models.go")
	}

	f.TemplateBody = modelsTemplate

	return nil
}

var modelsTemplate = `{{ .Boilerplate }}

package models

import (
	"context"
	"fmt"
	"{{ .Repo }}/pkg/ent"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}`
