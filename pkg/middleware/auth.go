package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ulyssesorz/douyin/pkg/jwt"
	"github.com/ulyssesorz/douyin/pkg/zap"
)

func TokenAuthMiddleware(jwt jwt.JWT, skipRoutes ...string) app.HandlerFunc {
	logger := zap.InitLogger()
	return func(ctx context.Context, c *app.RequestContext) {
		// 对于skip的路由不对他进行token鉴权
		for _, skipRoute := range skipRoutes {
			if skipRoute == c.FullPath() {
				c.Next(ctx)
				return
			}
		}

		// 获取token
		var token string
		if string(c.Request.Method()[:]) == "GET" {
			token = c.Query("token")
		} else if string(c.Request.Method()[:]) == "POST" {
			if strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
				token = c.PostForm("token")
			} else {
				token = c.Query("token")
			}
		} else {
			responseWithError(ctx, c, http.StatusBadRequest, "bad request")
			logger.Errorln("bad request")
			return
		}
		if token == "" {
			responseWithError(ctx, c, http.StatusUnauthorized, "token required")
			logger.Errorln("token required")
			return
		}

		claim, err := jwt.ParseToken(token)
		if err != nil {
			responseWithError(ctx, c, http.StatusUnauthorized, err.Error())
			logger.Errorln(err.Error())
			return
		}

		// 在上下文中向下游传递token
		c.Set("Token", token)
		c.Set("Id", claim.Id)

		c.Next(ctx) // 交给下游中间件
	}
}
