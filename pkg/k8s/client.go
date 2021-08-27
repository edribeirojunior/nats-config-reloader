package k8s

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func rConn() *kubernetes.Clientset {

	var err error
	var config *rest.Config

	if config, err = rest.InClusterConfig(); err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

func ReConn() *rest.Config {

	var err error

	var config *rest.Config
	if config, err = rest.InClusterConfig(); err != nil {
		log.Printf("error creating client configuration: %v\n", err)
	}

	return config
}
