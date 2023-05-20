package cmd

import "botanica/microservices/plants/processor"

func PlantsConfig() *processor.Config {
	return &plantsCfg
}
