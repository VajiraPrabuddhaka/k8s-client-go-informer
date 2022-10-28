package main

import (
	"context"
	"github.com/VajiraPrabuddhaka/k8s-client-go-informer/pkg/k8s/client"
	"github.com/VajiraPrabuddhaka/k8s-client-go-informer/pkg/k8s/httproute/gateway/clientset/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"log"
	gw_v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	"time"
)

func main() {

	stop := make(chan struct{})
	defer close(stop)
	//err = controller.Run(stop)
	c, _ := WatchResources(client.GetOutClusterClientSetV1alpha1())

	c.Run(stop)

	select {}

}

func WatchResources(clientSet v1alpha2.HttpRouteV1Alpha1Interface) (cache.Controller, cache.Store) {
	h := HttpRouteEventHandler{}
	httpRouteStore, httpRouteController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.HttpRoutes("default").List(context.TODO(), lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.HttpRoutes("default").Watch(context.TODO(), lo)
			},
		},
		&gw_v1alpha2.HTTPRoute{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    h.onAdd,
			UpdateFunc: h.OnUpdate,
			DeleteFunc: h.OnDelete,
		},
	)

	//go httpRouteController.Run(wait.NeverStop)
	return httpRouteController, httpRouteStore
}

// HttpRouteEventHandler is used to provide functions for resource event handler
type HttpRouteEventHandler struct {
}

func (h *HttpRouteEventHandler) onAdd(obj interface{}) {
	log.Printf("onAdd called : %v", obj)
}

func (h *HttpRouteEventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	log.Printf("onUpdate called, oldObj: %v newObj:%v", oldObj, newObj)
}

func (h *HttpRouteEventHandler) OnDelete(obj interface{}) {
	log.Printf("onAdd called : %v", obj)
}
