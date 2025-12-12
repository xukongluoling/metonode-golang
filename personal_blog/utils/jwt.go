package utils

import (
	"errors"
	"time"

	"metonode-golang/personal_blog/config"

	"github.com/golang-jwt/jwt/v5"
)

// 从配置文件获取JWT密钥
func getJWTSecret() []byte {
	return []byte(config.AppConfig.JWT.Secret)
}

// Claims 自定义JWT声明结构
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID uint, username string) (string, error) {
	// 从配置文件获取过期时间
	expireTime := time.Now().Add(time.Duration(config.AppConfig.JWT.ExpireHours) * time.Hour)

	// 创建自定义声明
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "personal_blog",
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取完整的编码后的字符串token
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证token并获取声明
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
