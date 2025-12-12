package utils

import (
	"context"
)

// 定义上下文键
type contextKey string

const (
	// UserIDKey 用户ID上下文键
	UserIDKey = contextKey("user_id")
	// UsernameKey 用户名上下文键
	UsernameKey = contextKey("username")
)

// SetUserInfo 将用户信息存储到请求上下文中
func SetUserInfo(ctx context.Context, userID uint, username string) context.Context {
	ctx = context.WithValue(ctx, UserIDKey, userID)
	ctx = context.WithValue(ctx, UsernameKey, username)
	return ctx
}

// GetUserIDFromContext 从上下文获取用户ID
func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(UserIDKey).(uint)
	return userID, ok
}

// GetUsernameFromContext 从上下文获取用户名
func GetUsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(UsernameKey).(string)
	return username, ok
}
