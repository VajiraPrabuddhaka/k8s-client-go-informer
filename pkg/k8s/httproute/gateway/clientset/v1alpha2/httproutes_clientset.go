package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
)

type HttpRouteV1Alpha1Interface interface {
	HttpRoutes(namespace string) HttpRouteInterface
}

type HttpRouteV1Alpha1Client struct {
	restClient rest.Interface
}

func (c *HttpRouteV1Alpha1Client) HttpRoutes(namespace string) HttpRouteInterface {
	return &httpRouteClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}

func NewForConfig(c *rest.Config) (*HttpRouteV1Alpha1Client, error) {
	v1alpha2.AddToScheme(scheme.Scheme)
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha2.GroupName, Version: v1alpha2.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &HttpRouteV1Alpha1Client{restClient: client}, nil
}
