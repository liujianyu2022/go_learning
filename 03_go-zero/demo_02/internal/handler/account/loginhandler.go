package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_learning/03_go-zero/demo_02/internal/logic/account"
	"go_learning/03_go-zero/demo_02/internal/svc"
	"go_learning/03_go-zero/demo_02/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
