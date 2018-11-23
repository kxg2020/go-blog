package controller

import (
	"github.com/gin-gonic/gin"
	"meeting/model"
)

func DepartmentInit(ctx *gin.Context)  {
	department := model.NewDepartment()
	result := department.GetDepartmentList()
	if result{
		PrintResult(9002,Empty,ctx)
	}
}