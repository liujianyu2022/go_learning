package db

import (
	"context"
	"go_learning/03_go-zero/demo_02/internal/config"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func NewMysql(mysqlConfig config.MySQLConfig) sqlx.SqlConn {
	mysql := sqlx.NewMysql(mysqlConfig.DataSource)
	
	db, err := mysql.RawDB()
	if err != nil {
		panic(err)
	}

	cxt, cancel := context.WithTimeout(
		context.Background(), 
		time.Second * time.Duration(mysqlConfig.ConnectTimeout),
	)

	defer cancel()
	err = db.PingContext(cxt)

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	
	return mysql
}

