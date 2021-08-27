package k8s

import (
	"context"
	"log"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

var decUnstructured = yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

func GetResource(ctx context.Context, cfg *rest.Config, s, ns string) (*unstructured.Unstructured, error) {
	log.Printf("Checking the Object Status...")

	dc, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return nil, err
	}

	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))

	client, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	obj := &unstructured.Unstructured{}
	_, gvk, err := decUnstructured.Decode([]byte(s), nil, obj)
	if err != nil {
		log.Printf("it was not possible initialize the GVK, following error happened: %s", err.Error())
		return nil, err
	}

	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		log.Printf("it was not possible initialize the GroupKind and Version, following error happened: %s", err.Error())
		return nil, err
	}

	unsTruct, err = client.Resource(mapping.Resource).Namespace(ns).Get(ctx, obj.GetName(), metav1.GetOptions{})
	if err != nil {
		log.Printf("It was not possible to get the object")
		return nil, err
	}

	if errors.IsNotFound(err) {
		return nil, err
	}

	return unsTruct, nil

}
