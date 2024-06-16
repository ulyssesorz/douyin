package zap

import (
	"github.com/ulyssesorz/douyin/pkg/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	config    = viper.Init("log")
	infoPath  = config.Viper.GetString("info")  //INFO&DEBUG&WARN级别的日志输出位置
	errorPath = config.Viper.GetString("error") //ERROR和FATAL级别的日志输出位置
)

func InitLogger() *zap.SugaredLogger {
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	encoder := getEncoder()
	//INFO级别的日志
	infoSyncer := getInfoWriter()
	infoCore := zapcore.NewCore(encoder, infoSyncer, lowPriority)

	//ERROR级别的日志
	errorSyncer := getErrorWriter()
	errorCore := zapcore.NewCore(encoder, errorSyncer, highPriority)

	//合并各种级别的日志
	core := zapcore.NewTee(infoCore, errorCore)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()
	return sugarLogger
}

// 定义日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getInfoWriter() zapcore.WriteSyncer {
	//lumberJack:日志切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   infoPath,
		MaxSize:    100, //日志文件最大100MB
		MaxBackups: 5,   //日志最多保存5个文件
		MaxAge:     30,  //最长保留时间
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getErrorWriter() zapcore.WriteSyncer {
	//lumberJack:日志切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   errorPath,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
