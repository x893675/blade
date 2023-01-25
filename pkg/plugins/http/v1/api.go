package v1

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/model/resource"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"

	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds"
)

// DefaultMainPath is default file path of main.go
const DefaultMainPath = "main.go"

var _ plugin.CreateAPISubcommand = &createAPISubcommand{}

type createAPISubcommand struct {
	config   config.Config
	resource *resource.Resource

	// force indicates that the resource should be created even if it already exists
	force bool

	// runMake indicates whether to run make or not after scaffolding APIs
	runMake bool
}

func (c *createAPISubcommand) UpdateMetadata(cliMeta plugin.CLIMetadata, subcmdMeta *plugin.SubcommandMetadata) {
	subcmdMeta.Description = `Scaffold a HTTP API by writing a Resource definition.

After the scaffold is written, the dependencies will be updated and
make generate will be run.
`
	subcmdMeta.Examples = fmt.Sprintf(`  # Create a frigates API with Group: ship, Version: v1beta1 and Kind: Frigate
  %[1]s create api --group ship --version v1beta1 --kind Frigate

  # Generate the swagger or other manifest
  make generate

  # build binary
  make build

  # run server direct
  make run
`, cliMeta.CommandName)
}

func (c *createAPISubcommand) BindFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&c.runMake, "make", true, "if true, run `make generate` after generating files")

	fs.BoolVar(&c.force, "force", false,
		"attempt to create resource even if it already exists")
}

func (c *createAPISubcommand) InjectConfig(cfg config.Config) error {
	c.config = cfg
	return nil
}

func (c *createAPISubcommand) InjectResource(r *resource.Resource) error {
	c.resource = r
	if err := c.resource.Validate(); err != nil {
		return err
	}
	// Check that resource doesn't have the API scaffolded or flag force was set
	if r, err := c.config.GetResource(c.resource.GVK); err == nil && r.HasAPI() && !c.force {
		return errors.New("API resource already exists")
	}
	return nil
}

func (c *createAPISubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := scaffolds.NewAPIScaffolder(c.config, *c.resource, c.force)
	scaffolder.InjectFS(fs)
	return scaffolder.Scaffold()
}

func (c *createAPISubcommand) PreScaffold(machinery.Filesystem) error {
	// check if main.go is present in the root directory
	if _, err := os.Stat(DefaultMainPath); os.IsNotExist(err) {
		return fmt.Errorf("%s file should present in the root directory", DefaultMainPath)
	}

	return nil
}

func (c *createAPISubcommand) PostScaffold() error {
	fmt.Print("Next: implement your new API.")
	return nil
}
