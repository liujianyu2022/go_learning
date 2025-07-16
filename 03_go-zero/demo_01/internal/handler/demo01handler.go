package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_learning/03_go-zero/demo_01/internal/logic"
	"go_learning/03_go-zero/demo_01/internal/svc"
	"go_learning/03_go-zero/demo_01/internal/types"
)

func Demo_01Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// 获取参数
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 注入serviceContext
		l := logic.NewDemo_01Logic(r.Context(), svcCtx)

		// 业务逻辑实现
		resp, err := l.Demo_01(&req)
		
		// 返回响应结果
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
