package model

import (
	"meeting/bootstrap"
)

// 获取用户
func QueryUserByUsername(username string)(map[string]interface{})  {
	user,err := bootstrap.Rose().Table("admin").Where("username","=",username).First()
	if err != nil{
		return ReturnResult(user,4000,err.Error())
	}
	if len(user) == 0{
		return ReturnResult(user,4000,"")
	}
	return ReturnResult(user,200,"")
}

// 更新用户
func UpdateUserAccount(data map[string]interface{},username string)map[string]interface{} {
	result,err := bootstrap.Rose().Table("admin").Where("username","=",username).Data(data).Update()
	if err != nil{
		return ReturnResult(result,500,err.Error())
	}
	return ReturnResult(result,200,"")
}
