package route

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const headerDefault  = ",Access-Control-Allow-Origin, Access-Control-Allow-Headers,Authorization,Content-Type"
const headerExpose   = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Content-Type,Authorization"
const originDefault  = "http://127.0.0.1:8080"
const allowedMethod  = "OPTIONS,POST,GET,PUT,DELETE"

// 跨域
func CrossSite()gin.HandlerFunc  {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		var headerString string
		for key,_ := range context.Request.Header{
			headerKeys = append(headerKeys,key)
		}
		headerString = strings.Join(headerKeys,",")
		if headerString == "" {
			headerString = headerDefault
		}else{
			headerString+= headerDefault
		}

		if origin == originDefault{
			// 允许请求的域
			context.Header("Access-Control-Allow-Origin",originDefault)
			// 允许请求的方法
			context.Header("Access-Control-Allow-Methods",allowedMethod)
			// 允许的header头
			context.Header("Access-Control-Allow-Headers",headerString)
			// options缓存时间
			context.Header("Access-Control-Max-Age","1800")
			// 返回的header头
			context.Header("Access-Control-Expose-Headers",headerExpose)
			// 是否验证cookie
			context.Header("Access-Control-Allow-Credentials","true")
			// 返回的数据缓存
			context.Header("Cache-Control","no-store")
		}
		if method == "OPTIONS"{
			context.AbortWithStatus(204)
		}else{
			context.Next()
		}
	}
}