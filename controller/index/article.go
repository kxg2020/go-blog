package index

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"go-blog/utils"
	"log"
	"time"
)

type Article struct {

}

func NewArticle()*Article  {
	return new(Article)
}

func (this *Article)GetArticleList(ctx *gin.Context)  {
	tag := ctx.PostForm("tag")
	result,err := model.GetArticleByTag(tag);
	if err != nil {
		log.Fatal(err.Error())
		utils.PrintErrors(4005,ctx);
		return
	}
	// 转换时间
	for _,value := range result{
		if val,ok := value["create_time"];ok && val != nil{
			value["create_time"] = time.Unix(val.(int64),0).Format("2006-01-02 15:04:05")
		}
	}
	utils.PrintSuccess(9005,result,ctx)
}

func (this *Article)GetArticleInfo(ctx *gin.Context)  {
	id := ctx.Param("id")
	result,err := model.GetArticleInfo(id)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4016,ctx)
		return
	}
	utils.PrintSuccess(9015,result,ctx)
}