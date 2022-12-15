package zoom_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	APIKey    string `yaml:"api_key"  mapstructure:"api_key"`
	APISecret string `yaml:"api_secret"  mapstructure:"api_secret"`
}
