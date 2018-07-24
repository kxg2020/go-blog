package model

import (
	"go-blog/db"
	"time"
)

type Tag struct {
	Id int `json:"id"`
	Tag_name string `json:"tag_name"`
	Create_time int `json:"create_time"`
	Status int  `json:"status"`
	Mark string `json:"mark"`
}

func GetTagList()([]map[string]interface{},error) {
	tag,err := db.Db().Table("tag").Get()
	if err != nil{
		return nil,err
	}
	for _ ,value := range tag{
		if val,ok := value["create_time"];ok && value["create_time"] != nil{
			value["create_time"] = time.Unix(val.(int64),0).Format("2006-01-02 15:04:05")
		}
	}
	return tag,nil
}

// 新增标签
func InsertTag(function func()map[string]interface{})(bool,error)  {
	insertData := function()
	result,err := db.Db().Table("tag").Data(insertData).Insert()
	if err != nil {
		return false,err
	}
	if result == 0 {
		return  false,nil
	}
	return true,nil
}

// 修改状态
func EditStatus(status string,id string) (bool,error) {
	updateData := map[string]interface{}{
		"status":status,
	}
	result,err := db.Db().Table("tag").Data(updateData).Where(map[string]interface{}{"id":id}).Update()
	if err != nil{
		return  false,err
	}
	if result > 0 {
		return true,nil
	}
	return  false,nil
}

// 保存信息
func EditTag(callback func()map[string]interface{},id string)(bool,error)  {
	updateData  := callback()
	result,err := db.Db().Table("tag").Where(map[string]interface{}{"id":id}).Data(updateData).Update()
	if err != nil{
		return  false,err
	}

	if result > 0 {
		return true,nil
	}
	return false,nil
}

// 删除标签
func DelTag(id string)(bool,error)  {
	result,err := db.Db().Table("tag").Where(map[string]interface{}{"id":id}).Delete()
	if err != nil{
		return false,err
	}
	if result > 0{
		return true,nil
	}
	return false,nil
}