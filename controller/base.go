package controller

import (
	"github.com/gin-gonic/gin"
)

var tips = map[int]string{
	4000:"用户名或密码错误",
	4001:"用户名或密码不能为空",
	4002:"修改用户信息失败",


	9000:"登陆成功",
	9001:"修改用户信息成功",
	9002:"部门初始化成功",
}
var Empty = make(map[string]interface{})

var result map[string]interface{}

func PrintResult(code int,data interface{},ctx *gin.Context){
	result = make(map[string]interface{})
	result["status"] = true
	if val,ok := tips[code];ok{
		result["msg"]  = val
	}
	if code <= 4000 && code != 200{
		result["status"] = false
	}
	result["data"] = data
	ctx.JSON(200,result)
}