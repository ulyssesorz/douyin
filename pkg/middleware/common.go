package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/gin-gonic/gin"
	z "github.com/ulyssesorz/douyin/pkg/zap"
	"go.uber.org/zap"
)

var (
	_      endpoint.Middleware = CommonMiddleware
	logger *zap.SugaredLogger
)

func init() {
	logger = z.InitLogger()
	defer logger.Sync()
}

func responseWithError(ctx context.Context, c *app.RequestContext, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"status_code": -1,
		"status_msg":  message,
	})
}

func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		logger.Debugf("real request: %+v", req)
		logger.Debugf("remote service name: %s, remote method: %s", ri.To().ServiceName(), ri.To().Method())
		if err := next(ctx, req, resp); err != nil {
			return err
		}
		logger.Infof("real response: %+v", resp)
		return nil
	}
}
