package bootstrap

import (
	"github.com/gin-gonic/gin"
	"wechat/pkg/bootstrap/middleware"
	"wechat/pkg/bootstrap/router"
)

type Console struct {
	gin *gin.Engine
}

// constructor
func NewConsole()*Console  {
	return new(Console)
}

// framework initialize
func (this *Console)FrameworkInit()*Console  {
	this.gin = middleware.Log(gin.New())
	this.gin = middleware.Recover(this.gin)
	return this
}

// register router
func (this *Console)RouterInit()*Console{
	this.gin = router.NewRouter(this.gin).RegisterRoute()
	return this
}


// start framework
func (this *Console)Run(port string)error  {
	err := this.gin.Run(port)
	if err != nil {
		return err
	}
	return nil
}

