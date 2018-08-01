package bootstrap

import (
	"github.com/gohouse/gorose"
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
)

// 保存数据库连接和gin服务器的结构体
type Boot struct {
	Connection gorose.Connection
	Router     *gin.Engine
	Redis	    redis.Conn
}

var BootInstance  = &Boot{}

// 初始化数据库和gin
func Init(options ...func(boot *Boot)) *Boot {
	for _,option := range options{
		option(BootInstance)
	}
	return BootInstance
}

// 获取数据库连接
func GetDb() *gorose.Database {
	return BootInstance.Connection.GetInstance()
}

// 获取redis连接
func GetRedis()redis.Conn  {
	return BootInstance.Redis
}
