/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"botanica/microservices/plants/processor"
	"os"

	"github.com/spf13/cobra"
)

var plantsCfg processor.Config

// plantsCmd represents the plants command
var plantsCmd = &cobra.Command{
	Use:   "plants",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		mongoURI, _ := cmd.Flags().GetString("mongo-uri")

		if len(mongoURI) == 0 {
			mongoURI = os.Getenv("MONGO")
		}
		plantsCfg = processor.Config{
			MongoURI: mongoURI,
		}

		plantsService := processor.NewtDefaultPlants(plantsCfg)
		plantsService.Start()
	},
}

func init() {
	rootCmd.AddCommand(plantsCmd)

	plantsCmd.PersistentFlags().String("mongo-uri", "", "URI for mongo connection")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// plantsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// plantsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
