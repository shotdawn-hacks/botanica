package processor

import (
	"botanica/microservices/plants/server"
	"botanica/pkg/destination"
)

type Plants struct {
	Address      string                     `json:"address" bson:"address"`
	Port         string                     `json:"port" bson:"port"`
	Destinations []*destination.Destination `json:"destinations" bson:"destinations"`
	Mongo        string
}

type Config struct {
	// TODO:
}

func NewtDefaultPlants(cfgs Config) *Plants {
	c := &Plants{
		Address: "",
		Port:    "9010",
	}

	return c
}

func (r *Plants) init() {
	for _, dest := range r.Destinations {
		dest.Start()
	}
}

func (r *Plants) Start() {
	r.init()

	api := server.NewRouter()

	api.Run(":" + r.Port)
}
