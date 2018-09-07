package backend

import (
	"backendApi/utils"
	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context)  {
	utils.PrintResult(ctx,9998,1,"")
}
