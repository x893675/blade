/*
Copyright 2022 The Kubernetes Authors.

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

package templates

import (
	"fmt"
	"strings"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &Readme{}

// Readme scaffolds a README.md file
type Readme struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.ProjectNameMixin

	License string
}

// SetTemplateDefaults implements file.Template
func (f *Readme) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "README.md"
	}

	f.License = strings.Replace(
		strings.Replace(f.Boilerplate, "/*", "", 1),
		"*/", "", 1)

	f.TemplateBody = fmt.Sprintf(readmeFileTemplate,
		codeFence("make build"))

	return nil
}

//nolint:lll
const readmeFileTemplate = `# {{ .ProjectName }}
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Build and Run
1. Install Instances of Custom Resources:

%s

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

## Refs

* [Echo Docs](https://echo.labstack.com/guide/customization/)
* [Swagger Generate](https://github.com/swaggo/swag)
* [Struct Validate](https://github.com/go-playground/validator)
* [Ent ORM](https://entgo.io/docs/getting-started)

## License
{{ .License }}
`

func codeFence(code string) string {
	return "```sh" + "\n" + code + "\n" + "```"
}
