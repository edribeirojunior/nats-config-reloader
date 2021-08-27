package aws

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go/aws"
)

func GetConfig(dnsName, hostedZoneName string) []string {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	svc := route53.NewFromConfig(cfg)

	pHZ := &route53.ListHostedZonesByNameInput{}

	hostedZone, _ := svc.ListHostedZonesByName(context.TODO(), pHZ)

	var recordSets []string
	for _, hzs := range hostedZone.HostedZones {
		if strings.Contains(*hzs.Name, hostedZoneName) {

			params := &route53.ListResourceRecordSetsInput{
				HostedZoneId:    aws.String(*hzs.Id),
				StartRecordName: aws.String(dnsName),
				StartRecordType: types.RRTypeA,
			}

			resp, err := svc.ListResourceRecordSets(context.TODO(), params)
			if err != nil {
				log.Fatalf("Not possible to list resources: %v", err)

			}

			for _, sets := range resp.ResourceRecordSets {
				if strings.Contains(*sets.Name, dnsName) && *&sets.Type == "A" {
					fmt.Println(*sets.Name, *&sets.Type)
					recordSets = append(recordSets, *sets.Name)
				}

			}

		}
	}

	return recordSets

}
