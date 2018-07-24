package model

import (
	"go-blog/db"
	"time"
)

type Article struct {
	Title      	  string `json:"title"   form:"title"`
	Tag_id        string `json:"tag"     form:"tag_id"`
	Status        string `json:"status"  form:"status"`
	Content       string `json:"content" form:"content"`
	Content_text  string `json:"content" form:"text"`
	Img_url       string `json:"img_url" form:"img_url"`
}

// 新增文章
func InsertArticle(article Article)(bool,error)  {
	insertData := map[string]interface{}{
		"title"  : article.Title,
		"content": article.Content,
		"tag_id" : article.Tag_id,
		"status" : article.Status,
		"create_time"  : time.Now().Unix(),
		"content_text" : article.Content_text,
		"img_url" : article.Img_url,
	}
	result,err := db.Db().Table("article").Data(insertData).Insert();
	if err != nil{
		return false,nil
	}
	if result >= 1{
		return true,nil
	}
	return false,nil
}

// 文章列表
func GetArticleList(pgNum int64,pgSize int64) ([]map[string]interface{},error) {
	Db := db.Db()
	start := (pgNum - 1) * pgSize
	result,err := Db.
		Fields("a.*,b.tag_name").
		Table("article  a").
		LeftJoin("xm_tag  b","a.tag_id","=","b.id").
		Limit(int(pgSize)).
		Offset(int(start)).
		Order("a.create_time desc").
		Get()
	if err != nil{
		return []map[string]interface{}{},err
	}
	return result,nil
}

// 文章总数
func GetArticleCount()(int,error)  {
	result,err := db.Db().
		Fields("a.id,a.title,a.content,a.create_time,b.tag_name,a.status").
		Table("article  a").
		LeftJoin("xm_tag  b","a.tag_id","=","b.id").
		Count()
	if err != nil {
		return 0,nil
	}
	return result,nil
}

// 编辑状态
func EditArticleStatus(id string,status string)(bool,error)  {
	updateData := map[string]interface{}{
		"status":status,
	}
	result,err := db.Db().Table("article").Where(map[string]interface{}{
		"id":id,
	}).Data(updateData).Update()
	if err != nil{
		return false,err
	}
	if result == 1{
		return true,nil
	}
	return  false,nil
}

// 删除文章
func DelArticle(id string)(bool,error)  {
	result,err := db.Db().Table("article").Where(map[string]interface{}{"id":id}).Delete()
	if err != nil {
		return  false,err
	}
	if result == 1 {
		return true,nil
	}
	return false,nil
}

// 文章详情
func GetArticleInfo(id string)(map[string]interface{},error)  {
	result,err := db.Db().
		Fields("a.*,b.tag_name").
		Table("article a").
		Where(map[string]interface{}{
		"a.id":id,
	}).
		LeftJoin("xm_tag b","a.tag_id","=","b.id").
		First();
	if err != nil{
		return map[string]interface{}{},err
	}

	return result,nil
}

// 保存修改
func SaveEdit(id string,article Article)(bool,error)  {
	updateData := map[string]interface{}{
		"title"       : article.Title,
		"tag_id"      : article.Tag_id,
		"status"      : article.Status,
		"content"     : article.Content,
		"content_text": article.Content_text,
		"img_url"     : article.Img_url,
	}
	result,err := db.Db().
		Table("article").
		Data(updateData).
		Where(map[string]interface{}{
		"id":id,
	}).Update()
	if err != nil {
		return false,err
	}
	if result >= 1{
		return  true,nil
	}
	return false,nil
}

// 标签筛选文章
func GetArticleByTag(tag string)([]map[string]interface{},error)  {
	Db := db.Db()
	if tag == "default"{
		tags,err := GetTagList()
		if err != nil{
			return []map[string]interface{}{},err
		}
		tag = tags[0]["tag_name"].(string);
	}
	result,err := Db.
		Table("tag a").
		Fields("b.*,a.tag_name").
		LeftJoin("xm_article b","a.id","=","b.tag_id").
		Where([][]interface{}{{"a.tag_name","=",tag}}).
		Get()

	if err != nil{
		return []map[string]interface{}{},nil
	}
	return result,nil
}

