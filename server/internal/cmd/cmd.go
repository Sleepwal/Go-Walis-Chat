package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"server/internal/controller"
	"server/internal/middleware"
	"server/internal/ws"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(ghttp.MiddlewareCORS)
				// 错误处理中间件
				group.Middleware(middleware.ErrorHandler)

				group.Bind(
					controller.Chat,
				)
			})

			go ws.HubObj.Run()
			s.Run()
			return nil
		},
	}
)
