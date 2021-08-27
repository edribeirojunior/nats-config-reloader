package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	aws "github.com/edribeirojunior/nats-config-reloader/pkg/aws"
	k8s "github.com/edribeirojunior/nats-config-reloader/pkg/k8s"
)

func ConfigReloader(dnsName, hostedZoneName, natsObj, natsNs string, timeout int) {
	cfg := k8s.ReConn()

	for {
		nServers := aws.GetConfig(dnsName, hostedZoneName)

		unstruct, err := k8s.GetResource(context.TODO(), cfg, natsObj, natsNs)

		if err != nil {
			log.Println(err)
		}

		natsObject := NatsCluster{}

		cltObj, _ := json.Marshal(unstruct)
		json.Unmarshal(cltObj, &natsObject)

		var superCluster []string

		for _, value := range natsObject.Spec.GatewayConfig.Gateways {
			superCluster = append(superCluster, value.Url)
		}

		if len(nServers) != len(superCluster) {
			gtwConfig := GtwConfiguration(nServers)
			natsObject.Spec.GatewayConfig.Gateways = gtwConfig

			newObject, _ := json.Marshal(natsObject)

			errUpdate := k8s.UpdateResource(context.TODO(), cfg, newObject)
			if errUpdate != nil {
				log.Println(errUpdate.Error())
			}

			if natsObject.Spec.Size > 1 {

				for i := 1; i < int(natsObject.Spec.Size); i++ {
					podName := fmt.Sprintf("%s-%i", natsObj, i)
					err := k8s.deletePod(podName, natsNs)
					if err != nil {
						log.Println(err.Error())
					}
				}
			}

		}

		time.Sleep(time.Second * timeout)

	}

}
