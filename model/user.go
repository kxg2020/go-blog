package model

import (
	"backendApi/bootstrap"
	"backendApi/service"
	"backendApi/utils"
	"strconv"
	"time"
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

func AddNewUser(user service.NewUser) (int,error) {
	saltNew     := utils.RandNumber(0,10000)
	passwordNew := utils.Md5Encrypt(user.Password + strconv.Itoa(int(saltNew)))
	statusNew   := func() int64{
		if user.Status{
			return 1
		}
		return  0
	}

	result,err  := bootstrap.GetDb().Table("user").Data(map[string]interface{}{
		"username" : user.Username,
		"password" : passwordNew,
		"salt"     : saltNew,
		"create_time" : time.Now().Unix(),
		"status"   : statusNew(),
	}).InsertGetId()
	if err != nil {
		return 0,err
	}
	if result != 0{
		return result,nil
	}
	return result,err
}

func GetUserList()([]map[string]interface{},error)  {
	user,err := bootstrap.GetDb().Fields("username,status,create_time,id").Table("user").Get()
	if err != nil {
		return []map[string]interface{}{},err
	}

	for _,value := range user{
		if val,ok := value["create_time"]; ok && val != "" {
			value["create_time"] = time.Unix(val.(int64),0).Format("2006-01-02 15:04:05")
		}
	}
	return user,nil
}

func DelUserById(id int)(int,error)  {
	result,err := bootstrap.GetDb().Table("user").Where(map[string]interface{}{"id":id}).Delete()
	if err != nil {
		return  0,err
	}
	return  result,nil
}

func ValidateLoginToken(username string,token string)bool  {
	user,err := bootstrap.GetDb().Table("user").
		Where([][]interface{}{{"username","=",username}}).
		Fields("token").
		First()
	if err != nil {
		return  false
	}
	if len(user) >= 1{
		if token == user["token"]{
			return true
		}else{
			return false
		}
	}
	return  false
}