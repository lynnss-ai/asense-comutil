package core

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

const (
	CtxAuth = "Authorization"
)

const (
	CtxAuthKeyByJwtUserID = "JwtUserID"
)

// GeneratorJwtToken Generator Token
func GeneratorJwtToken(val string, secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["iss"] = "lynntobing@gmail.com"
	claims[val] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetAuthJwtKeyValue[T any](ctx context.Context, val string) T {
	value := ctx.Value(val)
	return value.(T)
}

func GetAuthKeyJwtUserID(ctx context.Context) string {
	return GetAuthJwtKeyValue[string](ctx, CtxAuthKeyByJwtUserID)
}
