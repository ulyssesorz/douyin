// Package service /*
package service

import (
	"github.com/ulyssesorz/douyin/internal/tool"
	"github.com/ulyssesorz/douyin/pkg/jwt"
	"github.com/ulyssesorz/douyin/pkg/rabbitmq"
	"github.com/ulyssesorz/douyin/pkg/viper"
	"github.com/ulyssesorz/douyin/pkg/zap"
)

var (
	Jwt        *jwt.JWT
	logger     = zap.InitLogger()
	config     = viper.Init("rabbitmq")
	autoAck    = config.Viper.GetBool("consumer.relation.autoAck")
	RelationMq = rabbitmq.NewRabbitMQSimple("relation", autoAck)
	err        error
	privateKey string
)

func Init(signingKey string) {
	Jwt = jwt.NewJWT([]byte(signingKey))
	privateKey, _ = tool.ReadKeyFromFile(tool.PrivateKeyFilePath)

	go consume()
}
