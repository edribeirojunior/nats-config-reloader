package k8s

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeletePod(name, ns string) error {
	client := rConn()
	err := client.CoreV1().Pods(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}
