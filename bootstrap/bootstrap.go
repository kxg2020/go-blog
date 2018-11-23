package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
	_ "github.com/gohouse/gorose/driver/mysql"
	"log"
	"meeting/config"
)

var err error

type Bootstrap struct {
	// 引擎
	Router *gin.Engine
	// 数据库
	Rose   *gorose.Connection
}

// 启动全局对象
var Engine *Bootstrap

func NewBootstrap()*Bootstrap  {
	return new(Bootstrap)
}

// 初始化框架
func (this *Bootstrap)FrameInit(routerInit func(router *gin.Engine) *gin.Engine)*Bootstrap  {
	this.Router = routerInit(gin.New())
	return this
}

// 初始化数据库
func (this *Bootstrap)DbInit()*Bootstrap  {
	this.Rose,err = gorose.Open(config.Mysql)
	if err != nil {
		log.Fatal(err.Error())
		return  nil
	}
	return this
}


// 数据库连接
func Rose() *gorose.Session {
	return Engine.Rose.NewSession()
}



