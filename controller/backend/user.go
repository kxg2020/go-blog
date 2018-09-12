package backend

import (
	"backendApi/utils"
	"github.com/gin-gonic/gin"
	"log"
	"backendApi/service"
	"backendApi/model"
	"time"
	"strconv"
)

func GetUserList(ctx *gin.Context)  {
	user,err := model.GetUserList()
	if err != nil {
		log.Fatal(err);return
	}
	utils.PrintResult(ctx,9998,1,user)
}

func AddUser(ctx *gin.Context){
	var user service.NewUser
	if err := ctx.Bind(&user);err != nil{
		log.Fatal(err)
		return
	}
	result,err := model.AddNewUser(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	if result > 0{
		utils.PrintResult(ctx,9997,1,map[string]interface{}{
			"create_time" : time.Now().Format("2006-01-02 15:04:05"),
			"id"          : result,
		});return
	}
	if result < 0{
		utils.PrintResult(ctx,0010,0, "");return
	}
	utils.PrintResult(ctx,0004,0,"");return
}

func DelUser(ctx *gin.Context)  {
	id,_ := strconv.Atoi(ctx.PostForm("id"))
	result,err := model.DelUserById(id);
	if err != nil {
		log.Fatal(err)
		return
	}
	if result >= 1{
		utils.PrintResult(ctx,9997,1,"");return
	}
	utils.PrintResult(ctx,0006,0,"")
}

func EditUserStatus(ctx *gin.Context)  {
	id,_ := strconv.Atoi(ctx.PostForm("id"))
	status,_ := strconv.Atoi(ctx.PostForm("status"))
	result,err := model.EditUserStatus(id,status)
	if err != nil {
		log.Fatal(err)
		return
	}
	if result {
		utils.PrintResult(ctx,9995,1,"");return
	}
	utils.PrintResult(ctx,0012,1,"");return
}

func SearchUser(ctx *gin.Context)  {
	var search service.Search
	if err := ctx.Bind(&search);err != nil {
		log.Fatal(err)
		return
	}
	user,err := model.SearchUser(search)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(user) >= 1{
		utils.PrintResult(ctx,9994,1,user);return
	}
	utils.PrintResult(ctx,0014,0,user);return
}
