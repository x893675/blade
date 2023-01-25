module github.com/x893675/blade

go 1.19

require (
	github.com/onsi/ginkgo/v2 v2.5.1
	github.com/onsi/gomega v1.24.1
	github.com/spf13/afero v1.9.3
	github.com/spf13/pflag v1.0.5
	sigs.k8s.io/kubebuilder/v3 v3.8.0
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gobuffalo/flect v0.3.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/tools v0.3.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace sigs.k8s.io/kubebuilder/v3 => ./staging/src/sigs.k8s.io/kubebuilder
