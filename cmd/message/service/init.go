package service

import (
	"github.com/ulyssesorz/douyin/internal/tool"
	"github.com/ulyssesorz/douyin/pkg/jwt"
)

var (
	Jwt        *jwt.JWT
	publicKey  string
	privateKey string
)

func Init(signingKey string) {
	Jwt = jwt.NewJWT([]byte(signingKey))
	publicKey, _ = tool.ReadKeyFromFile(tool.PublicKeyFilePath)
	privateKey, _ = tool.ReadKeyFromFile(tool.PrivateKeyFilePath)
}
