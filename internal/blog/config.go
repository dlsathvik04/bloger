package blog

type Config struct{}

func NewConfig(configPath string) (*Config, error) {
	return &Config{}, nil
}
