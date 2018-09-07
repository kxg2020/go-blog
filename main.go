package main

import (
	"backendApi/bootstrap"
	"backendApi/middleware"
	"backendApi/controller/backend"
)

func main()  {
	boot := bootstrap.InitBootstrap()
	// 跨域
	boot.Router.Use(middleware.CrossSite())
	boot.Router.POST("/v1/login",backend.LoginValidate)
	// jwt
	boot.Router.Use(middleware.JwtAuth())
	boot.Router.POST("/v1/token",backend.TokenValidate)
	boot.Router.POST("/v1/userList",backend.GetUserList)
	boot.Router.Run(":8888")
}


