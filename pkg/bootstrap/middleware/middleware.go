package middleware

import (
	"github.com/gin-gonic/gin"
	"wechat/pkg/logs"
)

const options  = "OPTIONS"

// use log
func Log(framework *gin.Engine) *gin.Engine  {
	framework.Use(gin.Logger())
	return framework
}

// recover
func Recover(framework *gin.Engine)*gin.Engine  {
	callback := func()gin.HandlerFunc {
		return func(context *gin.Context) {
			defer func() {
				if err := recover(); err != nil {
					if err,ok := err.(string);ok{
						logs.Zap.Error(err)
					}
					context.JSON(500,map[string]interface{}{
						"code"   : "500",
						"msg"    : "inner error",
						"status" : false,
						"data"   : map[string]string{},
					})
					return
				}
			}()
			context.Next()
		}
	}
	framework.Use(callback())
	return framework
}

// cross site
func CrossSite(framework *gin.Engine)*gin.Engine  {
	framework.Use(func(context *gin.Context) {
		headers:= context.Request.Header.Get("Access-Control-Request-Headers")
		origin := context.Request.Header.Get("origin")
		method := context.Request.Method
		if origin == "http://127.0.0.1:8080" || origin == "http://127.0.0.1:8081" {
			// 允许请求的域
			context.Header("Access-Control-Allow-Origin", origin)
			// 允许请求的header头
			context.Header("Access-Control-Allow-Headers", headers)
			// 允许请求的方法类型
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			// 允许请求的缓存时间
			context.Header("Access-Control-Max-Age", "600")
			// 是否验证cookie
			context.Header("Access-Control-Allow-Credentials", "true")
			// 返回的数据内容是否缓存
			context.Header("Cache-Control", "no-store")
			// 返回的数据格式
			context.Set("Content-Type", "application/json")
		}
		if method == options{
			context.AbortWithStatus(204)
		}else{
			context.Next()
		}
	})
	return framework
}