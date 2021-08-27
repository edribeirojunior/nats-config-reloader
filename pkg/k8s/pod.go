package k8s

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func deletePod(name, ns string) error {
	client := rConn()
	err := client.CoreV1().Pods(ns).Delete(name, metav1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}
