package bootInject

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
	"go-blog/bootstrap"
)

var config = map[string]interface{} {
	"Default":"mysql_dev",
	"SetMaxOpenConns": 5,
	"SetMaxIdleConns": 2,
	"Connections":map[string]map[string]string{
		"mysql_dev": {
			"host": "127.0.0.1",
			"username": "root",
			"password": "root",
			"port": "3306",
			"database": "xm_blog",
			"charset": "utf8",
			"protocol": "tcp",
			"prefix": "xm_",
			"driver": "mysql",
		},
		"sqlite_dev": {
			"database": "./foo.db",
			"prefix": "",
			"driver": "sqlite3",
		},
	},
}

func BootDatabase() func(boot *bootstrap.Boot)  {
	return func(boot *bootstrap.Boot) {
		Connection,err := gorose.Open(config)
		if err != nil{
			panic(err.Error())
		}
		errs := Connection.Ping()
		if errs != nil{
			panic(errs)
		}
		boot.Connection = Connection
	}
}