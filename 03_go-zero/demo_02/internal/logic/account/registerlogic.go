package account

import (
	"context"
	"errors"
	"time"

	"go_learning/03_go-zero/demo_02/internal/model"
	"go_learning/03_go-zero/demo_02/internal/svc"
	"go_learning/03_go-zero/demo_02/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	userModel := model.NewUserModel(logic.svcCtx.SQLConnection)

	user, err := userModel.FindOneByUsername(logic.ctx, req.Username)
	if err != nil { 
		logic.Logger.Error("Failed to find user by username:", err)
		return nil, err
	}
	if user != nil {
		logic.Logger.Error("Username already exists")
		return nil, errors.New("此用户名已经注册")
	}

	_, err = userModel.Insert(logic.ctx, &model.User{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})

	if err != nil {
		logic.Logger.Info("User registered successfully")
		return nil, err
	}

	return
}
