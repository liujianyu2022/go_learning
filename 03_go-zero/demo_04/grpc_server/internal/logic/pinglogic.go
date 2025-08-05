package logic

import (
	"context"

	"go_learning/03_go-zero/demo_04/grpc_server/greet"
	"go_learning/03_go-zero/demo_04/grpc_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *greet.Request) (*greet.Response, error) {
	// todo: add your logic here and delete this line

	return &greet.Response{
		Pong: "pong",
	}, nil
}
