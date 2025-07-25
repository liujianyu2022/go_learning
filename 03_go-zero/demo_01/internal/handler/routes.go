// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.5

package handler

import (
	"net/http"

	"go_learning/03_go-zero/demo_01/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: Demo_01Handler(serverCtx),
			},
		},
	)
}
