package config_loader

func Load() Config {
	config := Config{}
	config.load()
	return config
}
