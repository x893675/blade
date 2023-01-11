module github.com/x893675/blade

go 1.19

require sigs.k8s.io/kubebuilder/v3 v3.8.0

require (
	github.com/gobuffalo/flect v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/tools v0.3.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace sigs.k8s.io/kubebuilder/v3  => ./staging/src/sigs.k8s.io/kubebuilder
