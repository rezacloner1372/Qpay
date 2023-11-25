package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type JWTConfig struct {
	Secret     string `koanf:"secret"`
	Expiration int    `koanf:"expiration_hours"`
}

func InitConfig() (*JWTConfig, error) {
	// Initialize koanf and load the configuration from sample-config.yaml
	var k = koanf.New(".")

	if err := k.Load(file.Provider("./sample-config.yaml"), yaml.Parser()); err != nil {
		panic(err)
	}

	var jwtConfig JWTConfig

	if err := k.UnmarshalWithConf("jwt", &jwtConfig, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
		panic(err)
	}
	return &jwtConfig, nil
}

func GenerateJWT(id uint) (string, error) {
	jwtConfig, err := InitConfig()
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

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	jwtConfig, err := InitConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check token expiration
	expirationTime, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid expiration time")
	}

	if float64(time.Now().Unix()) > expirationTime {
		return nil, fmt.Errorf("token is expired")
	}

	return claims, nil
}
