package svc

import (
	"go_learning/03_go-zero/demo_02/internal/config"
	"go_learning/03_go-zero/demo_02/internal/db"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	SQLConnection sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConnection := db.NewMysql(c.MySQLConfig)
	return &ServiceContext{
		Config: c,
		SQLConnection: sqlConnection,
	}
}
