package logger

type Config struct {
	Level       string `koanf:"level"`
	Encoding    string `koanf:"encoding"`
	Development bool   `koanf:"development"`
}

func DefaultConfig() *Config {
	return &Config{
		Development: true,
		Encoding:    "json",
		Level:       "info",
	}
}
