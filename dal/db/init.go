package db

import (
	"fmt"
	"time"

	"github.com/ulyssesorz/douyin/pkg/viper"
	"github.com/ulyssesorz/douyin/pkg/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	_db       *gorm.DB
	config    = viper.Init("db")
	zapLogger = zap.InitLogger()
)

func GetDB() *gorm.DB {
	return _db
}

// 获取数据库名
func getDsn(driverWithRole string) string {
	username := config.Viper.GetString(fmt.Sprintf("%s.username", driverWithRole))
	password := config.Viper.GetString(fmt.Sprintf("%s.password", driverWithRole))
	host := config.Viper.GetString(fmt.Sprintf("%s.host", driverWithRole))
	port := config.Viper.GetInt(fmt.Sprintf("%s.port", driverWithRole))
	Dbname := config.Viper.GetString(fmt.Sprintf("%s.database", driverWithRole))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)

	return dsn
}

func init() {
	dsn1 := getDsn("mysql.source")
	var err error
	_db, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err.Error())
	}
	dsn2 := getDsn("mysql.replica1")
	dsn3 := getDsn("mysql.replica2")
	_db.Use(dbresolver.Register(dbresolver.Config{
		// 1主，2 3从
		Sources:           []gorm.Dialector{mysql.Open(dsn1)},
		Replicas:          []gorm.Dialector{mysql.Open(dsn2), mysql.Open(dsn3)},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: false,
	}))

	if err := _db.AutoMigrate(&User{}, &Video{}, &Comment{}, &FavoriteVideoRelation{}, &FollowRelation{}, &Message{}, &FavoriteCommentRelation{}); err != nil {
		zapLogger.Fatalln(err.Error())
	}
	db, err := _db.DB()
	if err != nil {
		zapLogger.Fatalln(err.Error())
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
}
