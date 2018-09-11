package utils

import (
	"github.com/gin-gonic/gin"
)

var success  = map[int]string{
	9999 : "登陆成功",
	9998 : "获取用户列表",
	9997 : "添加用户成功",
	9996 : "删除用户成功",
}
var fail     = map[int]string{
	0000 : "用户名或密码不能为空",
	0002 : "用户名或密码错误",
	0004 : "添加用户失败",
	0006 : "删除用户失败",
}
var result map[string]interface{}

// 格式化输出
func PrintResult(ctx *gin.Context,code int,codeType int,data interface{}){
	result = make(map[string]interface{})
	switch codeType {
	case 1:
		result = map[string]interface{}{
			"msg"   : success[code],
			"status": codeType,
			"data"  : data,
		}
	case 0:
		result = map[string]interface{}{
			"msg"   : fail[code],
			"status": codeType,
			"data"  : data,
		}
	}
	ctx.JSON(200,result)
}



