package router

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func globalMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		func(ctx context.Context, req *app.RequestContext) {
			hlog.Debugf("req = %v, method = %s", req.FullPath(), req.Method())
		},
	}
}
