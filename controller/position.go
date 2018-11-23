package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"meeting/model"
)

func PositionList(ctx *gin.Context)  {
	instance := model.NewPosition()
	position,err := instance.PositionList()
	if err != nil {
		log.Fatal(err.Error())
		PrintResult(200,Empty,ctx)
		return
	}
	PrintResult(200,position,ctx)
}