package admin

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"io"
	"go-blog/utils"
	"time"
	"go-blog/model"
	"math"
	"strconv"
	"go-blog/bootstrap"
)

type Article struct {

}


type EditArticle struct {
	Title      string `json:"title"`
	Tag        string `json:"tag"`
	Status     string `json:"status"`
	Content    string `json:"content"`
	Id         string `json:"id"`
}

func NewArticle()*Article  {
	article := new(Article)
	return article
}

// 文章首页
func (article *Article)List(ctx *gin.Context)  {
	pgNum,_    := strconv.ParseInt(ctx.PostForm("pgNum"),10,64);
	pgSize,_   := strconv.ParseInt(ctx.PostForm("pgSize"),10,64)
	result,err := model.GetArticleList(pgNum,pgSize)
	if pgNum  < 1  {pgNum  = 1}
	if pgSize > 100{pgSize = 10}
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4005,ctx)
		return
	}
	for _,value   := range result{
		if val,ok := value["create_time"];ok && value["create_time"] != nil{
			value["create_time"] = time.Unix(val.(int64),0).Format("2006-01-02 15:04:05")
		}
		if _,ok := value["content_text"];ok && value["content_text"] != ""{
			// len("hello 世界")获取的是字符串的字节长度(中文unicode 2 字节,utf-8 3字节),如果要获取字符串的长度
			// 可以使用utf8.RuneCountInString(str)或者[]rune(str)
			textTemp := []rune(value["content_text"].(string))
			textLen  := len(textTemp)
			if textLen > 150{textLen = 150}
			value["content_text"] = string(textTemp[0 : textLen - 1]) + "..."
		}
	}
	count,err := model.GetArticleCount()
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4005,ctx)
		return
	}
	data := map[string]interface{}{
		"pages"  : math.Ceil(float64(len(result)) / (float64(pgSize))),
		"total"  : count,
		"article": result,
	}
	utils.PrintSuccess(9005,data,ctx)
	return
}

// 编辑文章
func (article *Article)Edit(ctx *gin.Context)  {
	if utils.IsAjax(ctx){
		var params EditArticle
		err := ctx.BindJSON(&params)
		if err != nil{
			log.Fatal(err.Error())
			utils.PrintErrors(4004,ctx)
			return
		}
	 	updateData := map[string]interface{}{
			"title":params.Title,
			"content":params.Content,
			"tag_id":params.Tag,
			"status":params.Status,
		}
		_,err = bootstrap.GetDb().Table("article").
			Data(updateData).
			Where(map[string]interface{}{"id":params.Id}).
			Update()
		if err != nil{
			log.Fatal(err.Error())
			utils.PrintErrors(4004,ctx);
			return
		}
		utils.PrintSuccess(9004,map[string]interface{}{},ctx)
		return
	}
	id := ctx.Query("id")
	articleInfo,err := bootstrap.GetDb().Table("article").Where(map[string]interface{}{"id":id}).First()
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	tag,err := bootstrap.GetDb().Table("tag").Where(map[string]interface{}{"status":1}).Get()
	if err != nil{
		log.Fatal(err.Error())
		return
	}

	ctx.HTML(200,"admin/article/edit.html",map[string]interface{}{
		"data":articleInfo,
		"tag":tag,
	})
}

// 新增文章
func (article *Article)Insert(ctx *gin.Context)  {
	var Article model.Article
	err := ctx.Bind(&Article)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4002,ctx)
		return
	}

	result,err := model.InsertArticle(Article)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4002,ctx)
		return
	}
	if result {
		utils.PrintSuccess(9002,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4002,ctx)
}

// 删除文章
func (article *Article)Delete(ctx *gin.Context){
	id := ctx.Param("id")
	result,err := model.DelArticle(id)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4003,ctx)
		return
	}
	if result{
		utils.PrintSuccess(9003,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4003,ctx)
}

// 编辑状态
func (article *Article)EditStatus(ctx *gin.Context)  {
	status := ctx.PostForm("status")
	id     := ctx.PostForm("id")
	result,err := model.EditArticleStatus(id,status)
	if err != nil {
		log.Fatal(err.Error())
		utils.PrintErrors(4015,ctx)
		return
	}
	if result{
		utils.PrintSuccess(9014,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4015,ctx)
}

// 文章详情
func (article *Article)ArticleInfo(ctx *gin.Context)  {
	id := ctx.Param("id")
	result,err := model.GetArticleInfo(id)
	if err != nil {
		log.Fatal(err.Error())
		utils.PrintErrors(4016,ctx)
		return
	}
	utils.PrintSuccess(9015,result,ctx)
}

// 保存修改
func (article *Article)SaveEdit(ctx *gin.Context) {
	var params model.Article
	err := ctx.Bind(&params)
	id  := ctx.PostForm("id")
	if err != nil {
		log.Fatal(err.Error())
		utils.PrintErrors(4004, ctx)
		return
	}
	result, err := model.SaveEdit(id, params)
	if result {
		utils.PrintSuccess(9004,map[string]interface{}{},ctx)
		return
	}
	utils.PrintErrors(4004, ctx)
}

// 上传图片
func (article *Article)Upload(ctx *gin.Context)  {
	file, header , err := ctx.Request.FormFile("upload")
	filename  := header.Filename
	date      := time.Now().Format("2006-01-02")
	dirPath   := "static/uploadFile/"+date
	filePath  := dirPath+"/"+filename
	exist,err := utils.PathExists(dirPath)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	if !exist{
		err = os.Mkdir(dirPath,os.ModePerm)
		if err != nil{
			log.Fatal(err.Error())
			return
		}
	}
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		utils.PrintErrors(4001,ctx)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		utils.PrintErrors(4001,ctx)
		return
	}
	// 回调函数
	callback := ctx.Query("CKEditorFuncNum")
	getImage := ctx.Query("backUrl")
	filePath  = "http://"+ctx.Request.Host+"/"+filePath
	// 普通上传成功返回的js
	//returnString := "<script>window.parent.CKEDITOR.tools.callFunction("+callback+",'"+filePath+"','上传成功');</script>"
	// 前后端分离涉及到跨域,跳转到前端的域,然后在前端域界面中执行js
	ctx.Redirect(301,getImage+"?ImageUrl="+filePath + "&Message=" +"success"+ "&CKEditorFuncNum="+ callback)
}

// 上传封面
func (article *Article)UploadCover(ctx *gin.Context)  {
	file, header , err := ctx.Request.FormFile("upload")
	filename  := header.Filename
	date      := time.Now().Format("2006-01-02")
	dirPath   := "static/uploadFile/"+date+"/cover"
	filePath  := dirPath+"/"+filename
	exist,err := utils.PathExists(dirPath)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	if !exist{
		err = os.MkdirAll(dirPath,os.ModePerm)
		if err != nil{
			log.Fatal(err.Error())
			return
		}
	}
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		utils.PrintErrors(4018,ctx)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		utils.PrintErrors(4018,ctx)
		return
	}
	filePath  = "http://"+ctx.Request.Host+"/"+filePath
	utils.PrintSuccess(9018,map[string]interface{}{
		"url":filePath,
	},ctx)
}


