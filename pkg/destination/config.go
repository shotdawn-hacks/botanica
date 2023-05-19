package destination

func NewConfig(t, address, port string) *Config {
	switch t {
	case TypePlants:
		return newPlantsConfig(address, port)
	default:
		return &Config{}
	}
}

func newPlantsConfig(address, port string) *Config {
	return &Config{Address: address, Port: port}
}
