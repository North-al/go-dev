package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	secretKey string
}

func NewJwt(secretKey string) *Jwt {
	return &Jwt{secretKey: secretKey}
}

// 生成token
func (j *Jwt) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"iss": "gin-web",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 过期时间24小时
		"sub": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// 解析token
func (j *Jwt) ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	userID, ok := claims["sub"].(int)
	if !ok {
		return 0, errors.New("invalid user ID")
	}

	return userID, nil
}
