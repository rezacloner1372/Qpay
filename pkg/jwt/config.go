package jwt

type Config struct {
	Secret     string `koanf:"secret"`
	Expiration int    `koanf:"expiration_hours"`
}
