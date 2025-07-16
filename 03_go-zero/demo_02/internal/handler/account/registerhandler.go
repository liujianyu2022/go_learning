package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_learning/03_go-zero/demo_02/internal/logic/account"
	"go_learning/03_go-zero/demo_02/internal/svc"
	"go_learning/03_go-zero/demo_02/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
