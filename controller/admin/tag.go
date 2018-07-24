package admin

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"log"
	"go-blog/utils"
	"time"
)

type Tag struct {

}
type EditTag struct {
	Id 			string `json:"id" form:"id"`
	Create_time string `json:"create_time" form:"create_time"`
	Mark 		string `json:"mark" form:"mark"`
	Tag_name 	string `json:"tag_name" form:"tag_name"`
	Status  	string `json:"status" form:"status"`
}
func NewTag() *Tag  {
	return new(Tag)
}

func (tag *Tag)GetTagList(ctx *gin.Context)  {
	result,err := model.GetTagList()
	if err != nil{
		log.Fatal(err)
		utils.PrintErrors(4006,ctx)
		return
	}
	utils.PrintSuccess(9006,result,ctx)
}

// 新增标签
func (tag *Tag)InsertTag(ctx *gin.Context)  {
	name := ctx.PostForm("name")
	status := ctx.PostForm("status")
	mark := ctx.PostForm("mark")
	result,err := model.InsertTag(func() map[string]interface{} {
		insertData := map[string]interface{}{
			"tag_name":name,
			"mark":mark,
			"status":status,
			"create_time":time.Now().Unix(),
		}
		return insertData
	})
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4011,ctx);
		return
	}
	if result{
		utils.PrintSuccess(9010,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4011,ctx)
}

// 修改状态
func (tag *Tag)EditTagStatus(ctx *gin.Context)  {
	id := ctx.Param("id")
	status := ctx.PostForm("status")
	result,err := model.EditStatus(status,id)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4012,ctx)
		return
	}
	if result{
		utils.PrintSuccess(9011,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4012,ctx)
}

// 修改标签
func (tag *Tag)EditTag(ctx *gin.Context)  {
	var params EditTag
	err := ctx.Bind(&params)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4013,ctx);
		return
	}
	result,err := model.EditTag(func() map[string]interface{} {
		timeResource,err := time.Parse("2006-01-02 15:04:05",params.Create_time)
		if err != nil{
			log.Fatal(err.Error())
			return nil
		}
		timeNew := timeResource.Unix()
		updateData := map[string]interface{}{
			"create_time": timeNew,
			"tag_name"   : params.Tag_name,
			"mark"       : params.Mark,
			"status"     : params.Status,
		}
		return  updateData
	},params.Id)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4013,ctx)
		return
	}
	if result{
		utils.PrintSuccess(9012,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4013,ctx)
}

// 删除标签
func (tag *Tag)DelTag(ctx *gin.Context)  {
	id := ctx.PostForm("id")
	result,err := model.DelTag(id)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4014,ctx)
		return
	}

	if result{
		utils.PrintSuccess(9013,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4014,ctx)
}