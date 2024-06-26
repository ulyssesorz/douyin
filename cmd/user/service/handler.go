package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ulyssesorz/douyin/dal/db"
	"github.com/ulyssesorz/douyin/internal/tool"
	"github.com/ulyssesorz/douyin/kitex/kitex_gen/user"
	"github.com/ulyssesorz/douyin/pkg/jwt"
	"github.com/ulyssesorz/douyin/pkg/minio"
	"github.com/ulyssesorz/douyin/pkg/snowflake"
	"github.com/ulyssesorz/douyin/pkg/zap"
)

type UserServiceImpl struct {}

func (s *UserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	logger := zap.InitLogger()

	// 查询用户是否已存在
	usr, err := db.GetUserByName(ctx, req.Username)
	if err != nil {
		logger.Errorln(err.Error())
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil
	} else if usr != nil {
		logger.Errorf("该用户名已存在：%s", usr.UserName)
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "该用户名已存在，请更换",
		}
		return res, nil
	}

	// 创建用户
	rand.Seed(time.Now().UnixMilli())
	usr = &db.User{
		ID: snowflake.GetID(),
		UserName: req.Username,
		Password: tool.Md5Encrypt(req.Password),
		Avatar: fmt.Sprintf("default%d.png", rand.Intn(10)),
	}
	if err := db.CreateUser(ctx, usr); err != nil {
		logger.Errorln(err.Error())
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil	
	}

	claims := jwt.CustomClaims{Id: int64(usr.ID)}
	claims.ExpiresAt = time.Now().Add(time.Minute * 5).Unix()
	token, err := Jwt.CreateToken(claims)
	if err != nil {
		logger.Errorf("发生错误：%v", err.Error())
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误：token 创建失败",
		}
		return res, nil
	}
	res := &user.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(usr.ID),
		Token:      token,
	}
	return res, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	logger := zap.InitLogger()

	usr, err := db.GetUserByName(ctx, req.Username)
	if err != nil {
		logger.Errorln(err.Error())
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "登录失败：服务器内部错误",
		}
		return res, nil
	} else if usr == nil {
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名不存在",
		}
		return res, nil
	}	

	// 比较密码
	if tool.Md5Encrypt(req.Password) != usr.Password {
		logger.Errorln("用户名或密码错误")
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
		}
		return res, nil
	}

	// 生成token，并返回结果
	claims := jwt.CustomClaims{
		Id: int64(usr.ID),
	}
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	token, err := Jwt.CreateToken(claims)
	if err != nil {
		logger.Errorf("发生错误：%v", err.Error())
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误：token 创建失败",
		}
		return res, nil
	}
	res := &user.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(usr.ID),
		Token:      token,
	}
	return res, nil
}

func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	logger := zap.InitLogger()
	userID := req.UserId

	usr, err := db.GetUserByID(ctx, userID)
	if err != nil {
		logger.Errorf("发生错误：%v", err.Error())
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误：获取背景图失败",
		}
		return res, nil
	} else if usr == nil {
		logger.Errorf("该用户不存在")
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "该用户不存在",
		}
		return res, nil
	}

	avatar, err := minio.GetFileTemporaryURL(minio.AvatarBucketName, usr.Avatar)
	if err != nil {
		logger.Errorf("Minio获取头像失败：%v", err.Error())
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误：获取头像失败",
		}
		return res, nil
	}
	backgroundImage, err := minio.GetFileTemporaryURL(minio.BackgroundImageBucketName, usr.BackgroundImage)
	if err != nil {
		logger.Errorf("Minio获取背景图失败：%v", err.Error())
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误：获取背景图失败",
		}
		return res, nil
	}

	res := &user.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User: &user.User{
			Id:              int64(usr.ID),
			Name:            usr.UserName,
			FollowCount:     int64(usr.FollowingCount),
			FollowerCount:   int64(usr.FollowerCount),
			IsFollow:        userID == int64(usr.ID),
			Avatar:          avatar,
			BackgroundImage: backgroundImage,
			Signature:       usr.Signature,
			TotalFavorited:  int64(usr.TotalFavorited),
			WorkCount:       int64(usr.WorkCount),
			FavoriteCount:   int64(usr.FavoriteCount),
		},
	}
	return res, nil
}