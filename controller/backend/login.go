package backend

import (
	"github.com/gin-gonic/gin"
	"backendApi/utils"
	"log"
	"backendApi/model"
	"time"
	"backendApi/middleware"
	"strconv"
)

func LoginValidate(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if len(username) == 0 || len(password) == 0{
		utils.PrintResult(ctx,0000,0,"")
		return
	}
	user,err := model.GetUserByUsername(username)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(user) >= 1{
		passwordDb := user["password"]
		saltDb     := user["salt"]
		passwordMd5:= utils.Md5Encrypt(password+saltDb.(string))
		if passwordDb == passwordMd5{
			saltNew:= utils.RandNumber(0,10000)
			passNew:= utils.Md5Encrypt(password + strconv.FormatInt(saltNew,10))
			timeNow:= time.Now().Unix()
			intNumber  := int(utils.RandNumber(1,100000))
			loginToken := utils.Md5Encrypt("user-login-token" + username + strconv.Itoa(intNumber))
			updateData := map[string]interface{}{
				"salt"     : saltNew,
				"password" : passNew,
				"token"    : loginToken,
				"last_login_time" : timeNow,
			}
			result,err := model.UpdateUserPasswordAndSalt(username,updateData)
			if err != nil{
				log.Fatal(err)
				return
			}
			if result {
				// 生成token
				token := middleware.SetToken(username);
				ctx.Header("user-token",token)
				ctx.Header("user-expire","2")
				ctx.Header("user-online",loginToken)
				utils.PrintResult(ctx,9999,1,map[string]string{"username":username})
			}else{
				utils.PrintResult(ctx,0002,0,"")
			}
		}else{
			utils.PrintResult(ctx,0002,0,"")
		}
	}else{
		utils.PrintResult(ctx,0002,0,"")
	}
}

func TokenValidate(ctx *gin.Context){
	utils.PrintResult(ctx,9999,1,"")
}



