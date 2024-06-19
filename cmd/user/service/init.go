package service

import (
	"github.com/ulyssesorz/douyin/pkg/jwt"
)

var (
	Jwt *jwt.JWT
)


func Init(signingKey string) {
	Jwt = jwt.NewJWT([]byte(signingKey))
}
