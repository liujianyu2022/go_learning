package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySQLConfig MySQLConfig
}

type MySQLConfig struct {
	DataSource     string
	ConnectTimeout int64
}