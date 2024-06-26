package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/ulyssesorz/douyin/cmd/api/rpc"
	"github.com/ulyssesorz/douyin/internal/response"
	"github.com/ulyssesorz/douyin/kitex/kitex_gen/user"
	kitex "github.com/ulyssesorz/douyin/kitex/kitex_gen/user"
)

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	//校验参数
	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, response.Register{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "用户名或密码不能为空",
			},
		})
		return
	}
	if len(username) > 32 || len(password) > 32 {
		c.JSON(http.StatusOK, response.Register{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "用户名或密码长度不能大于32个字符",
			},
		})
		return
	}
	req := &kitex.UserRegisterRequest{
		Username: username,
		Password: password,
	}
	res, _ := rpc.Register(ctx, req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusOK, response.Register{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.Register{
		Base: response.Base{
			StatusCode: 0,
			StatusMsg:  res.StatusMsg,
		},
		UserID: res.UserId,
		Token:  res.Token,
	})
}

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, response.Login{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "用户名或密码不能为空",
			},
		})
		return
	}
	req := &user.UserLoginRequest{
		Username: username,
		Password: password,
	}
	res, _ := rpc.Login(ctx, req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusOK, response.Login{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.Login{
		Base: response.Base{
			StatusCode: 0,
			StatusMsg:  res.StatusMsg,
		},
		UserID: res.UserId,
		Token:  res.Token,
	})
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	token := c.Query("token")
	if len(token) == 0 {
		c.JSON(http.StatusOK, response.UserInfo{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "token 已过期",
			},
			User: nil,
		})
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.UserInfo{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "user_id 不合法",
			},
			User: nil,
		})
		return
	}

	req := &kitex.UserInfoRequest{
		UserId: id,
		Token:  token,
	}
	res, _ := rpc.UserInfo(ctx, req)
	if res.StatusCode == -1 {
		c.JSON(http.StatusOK, response.UserInfo{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
			User: nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.UserInfo{
		Base: response.Base{
			StatusCode: 0,
			StatusMsg:  res.StatusMsg,
		},
		User: res.User,
	})
}
