package logic

import (
	"context"

	"go_learning/03_go-zero/demo_01/internal/svc"
	"go_learning/03_go-zero/demo_01/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Demo_01Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemo_01Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Demo_01Logic {
	return &Demo_01Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Demo_01Logic) Demo_01(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here
	resp = &types.Response{
		Message: "hello " + req.Name,
	}
	return
}
