package router

import (
	"github.com/gin-gonic/gin"
	"wechat/pkg/controller"
)

type Router struct {
	FrameWork *gin.Engine
}

func NewRouter(gin *gin.Engine)*Router  {
	router := new(Router)
	router.FrameWork = gin
	return router
}

func (this *Router)RegisterRoute()*gin.Engine  {
	this.FrameWork.Any("/wechat", controller.NewWeChat().EntryPoint)
	return this.FrameWork
}