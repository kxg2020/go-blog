package main

import (
	"github.com/gin-gonic/gin"
	"meeting/bootstrap"
	"meeting/controller"
	"meeting/route"
	"meeting/util"
)

func main()  {
	port := ":8888";
	bootstrap.Engine = bootstrap.NewBootstrap().FrameInit(func(router *gin.Engine) *gin.Engine {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		router.Use(route.CrossSite())
		router.POST("/login/check",controller.Check)
		router.Use(util.TokenValidate())
		{
			router.POST("/user/list",controller.UserList)
			router.POST("/user/update",controller.UserUpdate)
			router.GET("/position/list",controller.PositionList)
			router.POST("/token/validate",controller.TokenValidate)
			router.GET("/department/init",controller.DepartmentInit)
		}
		return router
	}).DbInit()

	bootstrap.Engine.Router.Run(port)

}