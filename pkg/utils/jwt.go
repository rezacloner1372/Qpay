package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type JWTConfig struct {
	Secret     string `koanf:"jwt.secret"`
	Expiration int    `koanf:"jwt.expiration_hours"`
}

var (
	k      = koanf.New(".")
	parser = yaml.Parser()
)

func InitConfig() {
	// Initialize koanf and load the configuration from sample-config.yaml
	if err := k.Load(file.Provider("sample-config.yaml"), parser); err != nil {
		panic(err)
	}
}

func GenerateJWT(id uint) (string, error) {
	var jwtConfig JWTConfig
	if err := k.Unmarshal("jwt", &jwtConfig); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtConfig.Expiration)).Unix()
	tokenString, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
