/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/x893675/blade/pkg/internal/validation"
	"github.com/x893675/blade/pkg/plugins/http/v1/scaffolds"

	"github.com/spf13/pflag"

	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin/util"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugins/golang"
)

// Variables and function to check Go version requirements.
var (
	goVerMin = golang.MustParse("go1.19.0")
	goVerMax = golang.MustParse("go2.0alpha1")
)

var _ plugin.InitSubcommand = &initSubcommand{}

type initSubcommand struct {
	config config.Config
	// For help text.
	commandName string

	// boilerplate options
	license string
	owner   string

	// go config options
	repo        string
	projectName string

	// flags
	fetchDeps          bool
	skipGoVersionCheck bool
}

func (p *initSubcommand) UpdateMetadata(cliMeta plugin.CLIMetadata, subcmdMeta *plugin.SubcommandMetadata) {
	p.commandName = cliMeta.CommandName

	subcmdMeta.Description = `Initialize a new project including the following files:
  - a "go.mod" with project dependencies
  - a "PROJECT" file that stores project configuration
  - a "Makefile" with several useful make targets for the project
  - a "main.go" file that creates the server
  - a "pkg" docs that include core files
`
	subcmdMeta.Examples = fmt.Sprintf(`  # Initialize a new project with your repo and name in copyright
  %[1]s init --plugins go/v1 --owner "Your name" --repo "Your Repo" --project-name "Your Project"
`, cliMeta.CommandName)
}

func (p *initSubcommand) BindFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&p.skipGoVersionCheck, "skip-go-version-check",
		false, "if specified, skip checking the Go version")

	// dependency args
	fs.BoolVar(&p.fetchDeps, "fetch-deps", true, "ensure dependencies are downloaded")

	// boilerplate args
	fs.StringVar(&p.license, "license", "apache2",
		"license to use to boilerplate, may be one of 'apache2', 'none'")
	fs.StringVar(&p.owner, "owner", "", "owner to add to the copyright")

	// project args
	fs.StringVar(&p.repo, "repo", "", "name to use for go module (e.g., github.com/user/repo), "+
		"defaults to the go package of the current working directory.")

	fs.StringVar(&p.projectName, "project-name", "", "name of this project")
}

func (p *initSubcommand) InjectConfig(c config.Config) error {
	p.config = c

	// Try to guess repository if flag is not set.
	if p.repo == "" {
		repoPath, err := golang.FindCurrentRepo()
		if err != nil {
			return fmt.Errorf("error finding current repository: %v", err)
		}
		p.repo = repoPath
	}

	// Assign a default project name
	if p.projectName == "" {
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %v", err)
		}
		p.projectName = strings.ToLower(filepath.Base(dir))
	}
	// Check if the project name is a valid k8s namespace (DNS 1123 label).
	if err := validation.IsDNS1123Label(p.projectName); err != nil {
		return fmt.Errorf("project name (%s) is invalid: %v", p.projectName, err)
	}
	if err := p.config.SetProjectName(p.projectName); err != nil {
		return err
	}

	return p.config.SetRepository(p.repo)
}

func (p *initSubcommand) PreScaffold(machinery.Filesystem) error {
	// Ensure Go version is in the allowed range if check not turned off.
	if !p.skipGoVersionCheck {
		if err := golang.ValidateGoVersion(goVerMin, goVerMax); err != nil {
			return err
		}
	}

	// Check if the current directory has not files or directories which does not allow to init the project
	return checkDir()
}

func (p *initSubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := scaffolds.NewInitScaffolder(p.config, p.license, p.owner)
	scaffolder.InjectFS(fs)
	err := scaffolder.Scaffold()
	if err != nil {
		return err
	}

	if !p.fetchDeps {
		fmt.Println("Skipping fetching dependencies.")
		return nil
	}

	err = util.RunCmd("Get Swagger generate", "make", "swagger")
	if err != nil {
		return err
	}

	return nil
}

func (p *initSubcommand) PostScaffold() error {
	if p.fetchDeps {
		err := util.RunCmd("Update swagger docs", "make", "generate")
		if err != nil {
			return err
		}
		err = util.RunCmd("Update dependencies", "go", "mod", "tidy")
		if err != nil {
			return err
		}
	}
	return nil
}

// checkDir will return error if the current directory has files which are not allowed.
// Note that, it is expected that the directory to scaffold the project is cleaned.
// Otherwise, it might face issues to do the scaffold.
func checkDir() error {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Allow directory trees starting with '.'
			if info.IsDir() && strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
				return filepath.SkipDir
			}
			// Allow files starting with '.'
			if strings.HasPrefix(info.Name(), ".") {
				return nil
			}
			// Allow files ending with '.md' extension
			if strings.HasSuffix(info.Name(), ".md") && !info.IsDir() {
				return nil
			}
			// Allow capitalized files except PROJECT
			isCapitalized := true
			for _, l := range info.Name() {
				if !unicode.IsUpper(l) {
					isCapitalized = false
					break
				}
			}
			if isCapitalized && info.Name() != "PROJECT" {
				return nil
			}
			// Allow files in the following list
			allowedFiles := []string{
				"go.mod", // user might run `go mod init` instead of providing the `--flag` at init
				"go.sum", // auto-generated file related to go.mod
			}
			for _, allowedFile := range allowedFiles {
				if info.Name() == allowedFile {
					return nil
				}
			}
			// Do not allow any other file
			return fmt.Errorf(
				"target directory is not empty (only %s, files and directories with the prefix \".\", "+
					"files with the suffix \".md\" or capitalized files name are allowed); "+
					"found existing file %q", strings.Join(allowedFiles, ", "), path)
		})
	if err != nil {
		return err
	}
	return nil
}
