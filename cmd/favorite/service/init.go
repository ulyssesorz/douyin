package service

import (
	"github.com/ulyssesorz/douyin/pkg/jwt"
	"github.com/ulyssesorz/douyin/pkg/rabbitmq"
	"github.com/ulyssesorz/douyin/pkg/viper"
	"github.com/ulyssesorz/douyin/pkg/zap"
)

var (
	Jwt        *jwt.JWT
	logger     = zap.InitLogger()
	config     = viper.Init("rabbitmq")
	autoAck    = config.Viper.GetBool("consumer.favorite.autoAck")
	FavoriteMq = rabbitmq.NewRabbitMQSimple("favorite", autoAck)
	err        error
)

func Init(signingKey string) {
	Jwt = jwt.NewJWT([]byte(signingKey))
	go consume()
}
