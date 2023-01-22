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
	cfgv1 "github.com/x893675/blade/pkg/config/v1"
	"github.com/x893675/blade/pkg/plugins/http"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
)

const pluginName = http.DefaultNameQualifier

var (
	pluginVersion            = plugin.Version{Number: 1}
	supportedProjectVersions = []config.Version{cfgv1.Version}
)

var _ plugin.Full = Plugin{}

// Plugin implements the plugin.Full interface
type Plugin struct {
	initSubcommand
	createAPISubcommand
}

// Name returns the name of the plugin
func (Plugin) Name() string { return pluginName }

// Version returns the version of the plugin
func (Plugin) Version() plugin.Version { return pluginVersion }

// SupportedProjectVersions returns an array with all project versions supported by the plugin
func (Plugin) SupportedProjectVersions() []config.Version { return supportedProjectVersions }

// GetInitSubcommand will return the subcommand which is responsible for initializing and common scaffolding
func (p Plugin) GetInitSubcommand() plugin.InitSubcommand { return &p.initSubcommand }

func (p Plugin) GetCreateAPISubcommand() plugin.CreateAPISubcommand { return &p.createAPISubcommand }
