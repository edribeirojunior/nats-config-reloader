package cmd

import (
	"fmt"
	"os"

	nc "github.com/edribeirojunior/nats-config-reloader/pkg/nats"
	"github.com/spf13/cobra"
)

var hostedZoneName, dnsName, natsObject, natsNamespace string
var natsTimeout int

var rootCmd = &cobra.Command{
	User:  "nats-config-reloader",
	Short: "Nats Config Reloader",
	Long:  "Nats Config Reloader is lightweight binary that will handle nats-cluster configuration",
	Run: func(cmd *cobra.Command, args []string) {

		nc.ConfigReloader(dnsName, hostedZoneName, natsObject, natsNamespace, natsTimeout)
	},
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVar(&hostedZoneName, "host-zone-name", "dev.outsystemscloudrd.net.", "Set the Hosted Zone Name")
	rootCmd.PersistentFlags().StringVar(&natsObject, "nats-name", "plat-dev-eu-west-1-01", "Set the nats CLuster name")
	rootCmd.PersistentFlags().StringVar(&dnsName, "dns-name", "nats.", "Set the nats Cluster name")
	rootCmd.PersistentFlags().StringVar(&natsNamespace, "namespace", "nats-io", "Set the nats Namespace")
	rootCmd.PersistentFlags().IntVar(&natsTimeout, "timeout", 30, "Set the Timeout of the Looping to check the Route53 Records")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
