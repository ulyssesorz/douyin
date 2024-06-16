package middleware

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ulyssesorz/douyin/pkg/zap"
)

// 限流中间件，使用令牌桶的方式处理请求
func TokenLimitMiddleware() app.HandlerFunc {
	logger := zap.InitLogger()

	return func(ctx context.Context, c *app.RequestContext) {
		token := c.GetString("Token")

		if !CurrentLimiter.Allow(token) {
			responseWithError(ctx, c, http.StatusForbidden, "request too fast")
			logger.Errorln("403: Request too fast.")
			return
		}
		c.Next(ctx)
	}
}
