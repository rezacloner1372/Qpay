package utils

import (
	"Qpay/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Config struct {
	Secret     string `koanf:"secret"`
	Expiration int    `koanf:"expiration_hours"`
}

var cfg *Config

func initConfig() {
	c, err := config.Load(false)
	if err != nil {
		panic(err)
	}
	cfg = c.JWT

}

func GenerateJWT(id uint) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(cfg.Expiration)).Unix()
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
