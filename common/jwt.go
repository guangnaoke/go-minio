package common

import (
	"minio_server/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("minio_key")

type Claims struct {
	UserID    int16
	Access    string
	AccessKey string
	Level     int
	jwt.StandardClaims
}

// 颁发token
func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserID:    user.UserID,    // ID
		Access:    user.Access,    // 权限
		AccessKey: user.AccessKey, // 账号
		Level:     user.Level,     // 等级
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "minio",
			Subject:   "token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
