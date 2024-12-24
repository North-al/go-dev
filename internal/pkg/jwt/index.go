package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"northal.com/config"
)

type Jwt struct {
}

func NewJwt() *Jwt {
	return &Jwt{}
}

// 生成token
func (j *Jwt) GenerateToken(userID int) (string, error) {

	fmt.Println(config.GetJwtConfig().TokenExpire, "config.GetJwtConfig().TokenExpire")
	fmt.Println(config.GetJwtConfig().SecretKey, "config.GetJwtConfig().TokenSecretKey")
	fmt.Println(time.Duration(config.GetJwtConfig().TokenExpire)*time.Hour, "time.Duration(config.GetJwtConfig().TokenExpire)*time.Hour")

	claims := jwt.MapClaims{
		"iss": "gin-web",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(), // 过期时间24小时
		"sub": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJwtConfig().SecretKey))
}

// 解析token
func (j *Jwt) ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtConfig().SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	userID, ok := claims["sub"]

	if !ok {
		return 0, errors.New("invalid user ID")
	}

	if id, ok := userID.(float64); ok {
		return int(id), nil
	}

	return 0, errors.New("invalid user ID type")
}
