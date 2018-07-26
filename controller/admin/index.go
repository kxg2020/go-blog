package admin

import (
	"github.com/gin-gonic/gin"
	"go-blog/utils"
)

type Index struct {

}

func NewIndex() * Index  {
	index := new(Index)
	return  index
}

func (index *Index)CheckToken(ctx *gin.Context)  {
	utils.PrintSuccess(9017,map[string]interface{}{},ctx)
}