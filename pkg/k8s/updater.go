package k8s

import (
	"context"
	"log"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/discovery"
	memory "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

//UpdateResources is a function to update if an object already exists
func UpdateResource(ctx context.Context, cfg *rest.Config, s []byte) error {

	//log.Printf("Trying to start with the object %s", s)
	dc, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return err
	}

	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))

	dyn, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return err
	}

	obj := &unstructured.Unstructured{}
	_, gvk, err := decUnstructured.Decode(s, nil, obj)
	if err != nil {
		log.Printf("it was not possible initialize the GVK, following error happened: %s", err.Error())
		return err
	}

	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		log.Printf("it was not possible initialize the GroupKind and Version, following error happened: %s", err.Error())
		return err
	}

	var dr dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		dr = dyn.Resource(mapping.Resource).Namespace(obj.GetNamespace())
	} else {
		dr = dyn.Resource(mapping.Resource)
	}

	_, err = dr.Update(context.TODO(), obj, metav1.UpdateOptions{})
	if err != nil {
		log.Printf("it was not possible to update the resource, following error happened: %s", err.Error())
	}

	return err
}
