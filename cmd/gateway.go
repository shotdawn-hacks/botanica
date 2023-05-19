/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"botanica/microservices/api-gateway/processor"
	"botanica/pkg/destination"

	"github.com/spf13/cobra"
)

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Starter for api-gateway microservice",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plantsAddress, _ := cmd.Flags().GetString("plants-address")
		plantsPort, _ := cmd.Flags().GetString("plants-port")
		plantsCfg := destination.NewConfig("plants", plantsAddress, plantsPort)

		gatewayCfg := processor.Config{
			Plants: plantsCfg,
		}

		gatewayService := processor.NewtDefaultGateway(gatewayCfg)

		gatewayService.Start()
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	gatewayCmd.PersistentFlags().String("plants-port", "", "Port of plants service")
	gatewayCmd.PersistentFlags().String("plants-address", "", "Address of plants service")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatewayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
