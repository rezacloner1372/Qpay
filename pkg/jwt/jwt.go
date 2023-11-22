package jwt

import (
	// "Qpay/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var cfg *Config

// func initConfig() {
// 	c, err := config.Load(false)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cfg = c.JWT

// }

func GenerateJWT(id uint, secret string, expiration int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expiration)).Unix()
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
