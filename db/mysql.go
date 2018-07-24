package db

import (
	"github.com/gohouse/gorose"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var config = map[string]interface{} {
	"Default":"mysql_dev",
	"SetMaxOpenConns": 0,
	"SetMaxIdleConns": 1,
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

func Db()*gorose.Database{
	connect,err := gorose.Open(config)
	if err != nil{
		log.Fatal(err.Error())
		return  nil
	}
	instance := connect.GetInstance()
	return  instance
}