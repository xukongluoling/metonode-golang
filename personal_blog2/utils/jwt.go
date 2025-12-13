package utils

import (
	"errors"
	"metonode-golang/personal_blog2/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 配置文件获取jwt配置
func getJWTSecret() []byte {
	return []byte(config.AppConfig.Jwt.Secret)
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(userID uint, username string) (string, error) {
	// 获取过期时间配置
	expire := time.Now().Add(time.Duration(config.AppConfig.Jwt.Expire) * time.Second)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "personal_blog2",
		},
	}
	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名获取完整的token字符串
	tokenStr, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析token获取用户信息
func ParseToken(tokenStr string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证token并获取信息
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
