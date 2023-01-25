package main

import (
	"log"

	"sigs.k8s.io/kubebuilder/v3/pkg/cli"

	cfgv1 "github.com/x893675/blade/pkg/config/v1"
	httpv1 "github.com/x893675/blade/pkg/plugins/http/v1"
)

func main() {
	c, err := cli.New(
		cli.WithCommandName("blade"),
		cli.WithDescription("CLI tool for building HTTP/GRPC Server."),
		cli.WithVersion(versionString()),
		cli.WithPlugins(
			httpv1.Plugin{},
		),
		cli.WithDefaultPlugins(cfgv1.Version, httpv1.Plugin{}),
		cli.WithDefaultProjectVersion(cfgv1.Version),
		//cli.WithExtraCommands(commands...),
		//cli.WithExtraAlphaCommands(alphaCommands...),
		cli.WithCompletion(),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
