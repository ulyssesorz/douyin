package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var (
	_ endpoint.Middleware = ServerMiddleware
)

func ServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		logger.Infof("client address: %v", ri.From().Address())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}
