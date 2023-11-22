package logger

type Config struct {
	Level       string `koanf:"level"`
	Encoding    string `koanf:"encoding"`
	Development bool   `koanf:"development"`
}
