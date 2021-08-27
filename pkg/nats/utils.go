package nats

import "strings"

func GtwConfiguration(natsCluster []string) []Gateway {
	var gtwCfg []Gateway
	for _, i := range natsCluster {
		splitNats := strings.Split(i, ".")
		gtw := Gateway{Name: splitNats[1], Url: i}
		gtwCfg = append(gtwCfg, gtw)
	}

	return gtwCfg
}
