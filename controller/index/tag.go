package index

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"log"
	"go-blog/utils"
)

type Tag struct {

}

func NewTag()*Tag  {
	return  new(Tag)
}

func (tag *Tag)GetTagList(ctx *gin.Context)  {
	tags,err := model.GetTagList()
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4006,nil)
		return
	}
	if len(tags) > 0 {
		utils.PrintSuccess(9006,tags,ctx)
		return
	}
	utils.PrintErrors(4006,nil)
}