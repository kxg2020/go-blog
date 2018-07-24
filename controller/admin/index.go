package admin

import (
	"github.com/gin-gonic/gin"
)

type Index struct {

}

func NewIndex() * Index  {
	index := new(Index)
	return  index
}

func (index *Index)Index(ctx *gin.Context)  {

}
func (index *Index)CheckToken(ctx *gin.Context)  {

}