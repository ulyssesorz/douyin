package service

import (
	"github.com/ulyssesorz/douyin/pkg/jwt"
)

var (
	Jwt *jwt.Jwt
)

func init(signingKey string) {
	Jwt = jwt.NewJWT([]byte(signingKey))
}
