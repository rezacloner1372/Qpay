package db

type Config struct {
	Host     string `koanf:"mysql_host"`
	Port     int    `koanf:"mysql_port"`
	Username string `koanf:"mysql_username"`
	Password string `koanf:"mysql_password"`
	Database string `koanf:"mysql_database"`
}
