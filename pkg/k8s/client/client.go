package client

import (
	"flag"
	"github.com/VajiraPrabuddhaka/k8s-client-go-informer/pkg/k8s/httproute/gateway/clientset/v1alpha2"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var kubeconfig *string

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
}
func GetInClusterClientSet() *kubernetes.Clientset {
	var clientset *kubernetes.Clientset
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	clientset = c
	return clientset
}

func GetOutClusterClientSetKubernetes() *kubernetes.Clientset {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)

	//TODO: handle error
	return clientset
}

func GetOutClusterClientSetV1alpha1() *v1alpha2.HttpRouteV1Alpha1Client {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := v1alpha2.NewForConfig(config)

	//TODO: handle error
	return clientset
}
