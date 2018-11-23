package controller

import (
	"github.com/gin-gonic/gin"
	"meeting/model"
	"meeting/util"
	"strconv"
	"time"
)

// 用户登陆
func Check(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if util.CheckLength(username) && util.CheckLength(password){
		result := model.QueryUserByUsername(username)
		if ok := result["status"].(bool);ok{
			salt := result["data"].(map[string]interface{})["salt"].(int64)
			passwordSalt := util.Md5Encrypt(password + strconv.Itoa(int(salt)))
			passwordDb   := result["data"].(map[string]interface{})["password"].(string)
			if passwordSalt == passwordDb{
				saltNew := strconv.FormatInt(util.RandNumber(0,9999),10)
				passNew := util.Md5Encrypt(password + saltNew)
				update  := map[string]interface{}{
					"salt"     : saltNew,
					"password" : passNew,
					"last_login_time":time.Now().Unix(),
				}
				result  := model.UpdateUserAccount(update,username)
				if val := result["data"].(int64);val >= 1{
					token := util.CreateToken("meeting",username)
					ctx.Header("Authorization",token)
					PrintResult(9000,map[string]string{"username":username},ctx);return
				}else{
					PrintResult(4000,Empty,ctx);return
				}
			}else{
				PrintResult(4000,Empty,ctx);return
			}
		}else{
			PrintResult(4000,Empty,ctx);return
		}
	}
	PrintResult(4001,Empty,ctx)
}

func TokenValidate(ctx *gin.Context)  {
	ctx.JSON(200,map[string]interface{}{
		"status" : 1,
		"msg"    : "",
	})
	return
}

