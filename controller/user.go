package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"meeting/model"
	"meeting/util"
	"strconv"
)

func UserList(ctx *gin.Context)  {
	user := model.NewUser()
	var params model.User
	if err := ctx.BindJSON(&params);err != nil{
		log.Fatal(err.Error())
		PrintResult(200,Empty,ctx)
		return
	}
	pgNum,_ := strconv.Atoi(ctx.Query("pgNum"))
	pgSize,_:= strconv.Atoi(ctx.Query("pgSize"))
	pgNum,pgSize = util.PageFormat(pgNum,pgSize)
	users,count,err := user.UserList(pgNum,pgSize,params)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	PrintResult(200,map[string]interface{}{"user":users,"count":count},ctx)
}

func UserUpdate(ctx *gin.Context)  {
	var params model.User
	user := model.NewUser()
	if err := ctx.BindJSON(&params);err != nil{
		log.Fatal(err)
		return
	}
	result,err := user.UserUpdate(params)
	if err != nil {
		log.Fatal(err.Error())
		PrintResult(4002,Empty,ctx)
		return
	}
	if result >= 1{
		PrintResult(9001,Empty,ctx)
	}
}
