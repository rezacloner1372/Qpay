package server

type Config struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
}
