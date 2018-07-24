package utils

import (
	"github.com/gin-gonic/gin"
)

var Errors = map[int]string{
	4000:"用户名或密码错误",
	4001:"上传文件失败",
	4002:"添加文章失败",
	4003:"删除文章失败",
	4004:"更新文章失败",
	4005:"获取文章列表失败",
	4006:"获取标签列表失败",
	4007:"token已失效,请重新登陆",
	4008:"修改用户状态失败",
	4009:"修改用户信息失败",
	4010:"删除用户失败",
	4011:"添加标签失败",
	4012:"修改标签状态失败",
	4013:"修改标签信息失败",
	4014:"删除标签失败",
	4015:"修改文章状态失败",
	4016:"获取文章详情失败",
	4017:"添加用户失败",
	4018:"封面图上传成功",
}

var Success = map[int]string{
	9000:"登陆成功",
	9001:"上传文件成功",
	9002:"添加文章成功",
	9003:"删除文章成功",
	9004:"更新文章成功",
	9005:"获取文章列表成功",
	9006:"获取标签列表成功",
	9007:"修改用户状态成功",
	9008:"修改用户信息成功",
	9009:"删除用户成功",
	9010:"添加标签成功",
	9011:"修改标签状态成功",
	9012:"修改标签信息成功",
	9013:"删除标签成功",
	9014:"修改文章状态成功",
	9015:"获取文章详情成功",
	9016:"添加用户成功",
	9017:"token验证通过",
	9018:"封面图上传成功",
}

type PrintFormat struct {

}

func PrintSuccess(code int,data interface{},ctx *gin.Context){
	returnData := map[string]interface{}{
		"status": 1,
		"msg"   : Success[code],
		"data"  : data,
	}
	ctx.JSON(200,returnData)
}

func PrintErrors(code int,ctx *gin.Context) {
	returnData := map[string]interface{}{
		"status" : 0,
		"msg"    : Errors[code],
	}
	ctx.JSON(200,returnData)
}

func PrintTokenExpire(code int,ctx *gin.Context)  {
	returnData := map[string]interface{}{
		"status" : 2,
		"msg"    : Errors[code],
	}
	ctx.JSON(200,returnData)
}