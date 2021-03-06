module temp/kubebuild

go 1.13

require (
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/go-openapi/loads v0.19.4
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190723091251-e0797f438f94 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kubernetes-incubator/reference-docs v0.0.0 // indirect
	github.com/markbates/inflect v0.0.0-00010101000000-000000000000 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/pkg/errors v0.8.1 // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	k8s.io/api v0.18.4 // indirect
	k8s.io/apimachinery v0.18.4
	k8s.io/apiserver v0.18.4
	k8s.io/client-go v0.18.4
	k8s.io/code-generator v0.18.4 // indirect
	k8s.io/klog v1.0.0
	k8s.io/kube-aggregator v0.18.4 // indirect
	sigs.k8s.io/apiserver-builder-alpha v1.18.0
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/controller-tools v0.1.12 // indirect
	sigs.k8s.io/kubebuilder v1.0.8 // indirect
	sigs.k8s.io/structured-merge-diff v0.0.0-20190525122527-15d366b2352e // indirect
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
)

replace sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.12

replace sigs.k8s.io/kubebuilder => sigs.k8s.io/kubebuilder v1.0.8

replace github.com/markbates/inflect => github.com/markbates/inflect v1.0.4

replace github.com/kubernetes-incubator/reference-docs => github.com/kubernetes-sigs/reference-docs v0.0.0-20170929004150-fcf65347b256

replace sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v1.0.1-0.20191108220359-b1b620dd3f06
