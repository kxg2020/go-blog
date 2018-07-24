package admin

import (
	"go-blog/model"
	"log"
	"github.com/gin-gonic/gin"
	"go-blog/utils"
	"time"
	"strconv"
)

type User struct {

}

func NewUser() *User  {
	return  new(User)
}

func (user *User)GetUserList(ctx *gin.Context)  {
	result,err := model.GetUserList()
	if err != nil{
		log.Fatal(err)
		utils.PrintErrors(4005,ctx)
		return
	}
	utils.PrintSuccess(9005,result,ctx)
}

func (user *User)EditUserStatus(ctx *gin.Context)  {
	userId := ctx.PostForm("id")
	status := ctx.PostForm("status")
	if status == "true"{status = "1"}
	if status == "false"{status = "0"}
	result,err := model.UpdateUserStatus(status,userId)
	if err != nil && result{
		log.Fatal(err.Error())
		utils.PrintErrors(4008,ctx)
		return ;
	}
	utils.PrintSuccess(9007,map[string]interface{}{},ctx)
}

func (user *User)SaveUseEdit(ctx *gin.Context) {
	username   := ctx.PostForm("username")
	createTime := ctx.PostForm("create_time")
	status     := ctx.PostForm("status")
	userId     := ctx.PostForm("id")
	password   := ctx.PostForm("password")
	passwordOld := ctx.PostForm("passwordOld")

	if status == "true"{status = "1"}else{status = "0"}
	createTimeTemp,_ := time.Parse("2006-01-02 15:04:05",createTime)
	createTimeNew := createTimeTemp.Unix()
	// 更新盐和密码
	saltNew     := utils.RandInt64(0,999)
	passwordNew := utils.NewEncrypt().Md5(password + strconv.FormatInt(saltNew,10))

	updateData := map[string]interface{}{
		"status":status,
		"username":username,
		"create_time":createTimeNew,
	}
	if password != passwordOld{
		updateData["password"] = passwordNew
		updateData["salt"]     = saltNew
	}
	result,err := model.SaveUserEdit(updateData,userId)
	if result{
		utils.PrintSuccess(9008,map[string]interface{}{},ctx)
		return
	}
	log.Fatal(err.Error())
	utils.PrintErrors(4009,ctx)
}

func (user *User)DelUser(ctx *gin.Context)  {
	id := ctx.PostForm("id")
	result,err := model.DelUser(id)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4010,ctx)
		return
	}
	if result{
		utils.PrintSuccess(9009,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4010,ctx)
}

func (user *User)Insert(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	status   := ctx.PostForm("status")
	saltNew  := utils.RandInt64(0,999)
	passwordNew := utils.NewEncrypt().Md5(password + strconv.FormatInt(saltNew,10))
	insertData  := map[string]interface{}{
		"username" : username,
		"password" : passwordNew,
		"status"   : status,
		"salt"     : saltNew,
	}
	result,err := model.InsertUser(insertData)
	if err != nil {
		log.Fatal(err.Error())
		utils.PrintErrors(4017,ctx)
		return
	}
	if result {
		utils.PrintSuccess(9016,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4017,ctx)
}
