package config

import "github.com/gohouse/gorose"

// 数据库配置
var Mysql = &gorose.DbConfigSingle{
	Driver         : "mysql",
	EnableQueryLog :  true,
	SetMaxOpenConns: 20,
	SetMaxIdleConns: 1,
	Prefix         : "kxg_",
	Dsn            : "root:root@tcp(127.0.0.1:3306)/kxg_meeting?charset=utf8",
}



