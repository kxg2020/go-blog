package model

import (
	"time"
	"go-blog/bootstrap"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 获取用户列表
func GetUserList()([]map[string]interface{},error)  {
	user,err := bootstrap.GetDb().
		Fields("username,create_time,id,last_login_time,status,password").
		Table("user").Get()
	if err != nil{
		return nil,err
	}
	for _,value := range user{
		if val,ok := value["create_time"];ok{
			value["create_time"] = time.Unix(val.(int64),0).Format("2006-01-02 15:04:05")
		}
		if val,ok := value["last_login_time"];ok{
			value["last_login_time"] = time.Unix(val.(int64),0).Format("2006-01-02 15:04:05")
		}
	}
	return user,nil
}

// 获取用户详情
func GetUserInfo(username string)(map[string]interface{},error)  {
	user,err := bootstrap.GetDb().
		Fields("salt,password,id,username").
		Table("user").
		Where("username","=",username).
		First()
	if err != nil{
		return nil,err
	}
	return user,nil
}

// 更新用户密码
func UpdateUserPassword(salt int64,password string,id int64)(bool)  {
	updateData := map[string]interface{}{
		"salt"    :salt,
		"password":password,
		"last_login_time": time.Now().Unix(),
	}
	result,err := bootstrap.GetDb().Table("user").Where("id","=",id).Data(updateData).Update()
	if err != nil && result > 0  {
		return  false
	}
	return  true
}

// 更新用户状态
func UpdateUserStatus(status string,userId string)(bool,error)  {

	result,err := bootstrap.GetDb().Table("user").Data(map[string]interface{}{
		"status":status,
	}).Where("id","=",userId).Update()

	if err != nil || result <= 0{
		return  false,nil
	}
	return true,nil
}

// 保存用户
func SaveUserEdit(updateData map[string]interface{},userId string)(bool,error)  {
	_,err := bootstrap.GetDb().
		Table("user").
		Data(updateData).Where(map[string]interface{}{
			"id":userId,
	}).Update()
	if err != nil{
		return false,err
	}
	return true,nil
}

// 删除用户
func DelUser(id string)(bool,error)  {
	result,err := bootstrap.GetDb().Table("user").Where(map[string]interface{}{"id":id}).Delete()
	if err != nil{
		return false,err
	}
	if result > 0 {
		return true,nil
	}
	return  false,nil
}

// 添加用户
func InsertUser(params map[string]interface{})(bool,error)  {
	insertData := map[string]interface{}{
		"username"    : params["username"],
		"password"    : params["password"],
		"salt"        : params["salt"],
		"create_time" : time.Now().Unix(),
		"status"      : params["status"],
	}
	result,err := bootstrap.GetDb().Table("user").Data(insertData).Insert();
	if err != nil {
		return false,err
	}
	if result >= 1 {
		return  true,nil
	}
	return  false,nil
}