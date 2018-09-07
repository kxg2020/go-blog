package model

import (
	"backendApi/bootstrap"
	"fmt"
)

func GetUserByUsername(username string)(map[string]interface{},error)  {
	fields := "username,salt,id,password"
	user,err := bootstrap.GetDb().Table("user").Fields(fields).Where("username","=",username).First()
	if err != nil {
		return map[string]interface{}{},nil
	}
	return user,nil
}

func UpdateUserPasswordAndSalt(username string,data map[string]interface{})(bool,error)  {
	result,err := bootstrap.GetDb().Table("user").Where("username","=",username).Data(data).Update()
	if err != nil {
		return false,err
	}
	if result >= 0{
		return true,nil
	}
	return false,nil
}

func ValidateLoginToken(username string,token string)bool  {
	user,err := bootstrap.GetDb().Table("user").
		Where([][]interface{}{{"username","=",username}}).
		Fields("token").
		First()
	if err != nil {
		return  false
	}
	fmt.Println(token)
	fmt.Println(user["token"])
	if len(user) >= 1{
		if token == user["token"]{
			return true
		}else{
			return false
		}
	}
	return  false
}