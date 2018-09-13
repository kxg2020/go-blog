package model

import (
	"backendApi/bootstrap"
	"backendApi/service"
	"backendApi/utils"
	"strconv"
	"time"
	"log"
	"reflect"
	"strings"
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
	userExist,err := GetUserByUsername(user.Username)
	if err != nil {
		log.Fatal(err)
		return 0,err
	}
	if len(userExist) >= 1{
		return -1,nil
	}
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

func GetUserList(search service.Search)([]map[string]interface{},error,int)  {
	start := (search.Page - 1) * search.Size
	where := SearchUserCondition(search)
	db := bootstrap.GetDb()
	user,err := db.
		Fields("username,status,create_time,id").
		Table("user").
		Where(where).
		Limit(search.Size).
		Offset(start).
		Order("create_time desc").
		Get()
	count,err := bootstrap.GetDb().Where(where).Table("user").Count("id")
	if err != nil {
		return []map[string]interface{}{},err,0
	}
	for _,value := range user{
		if val,ok := value["create_time"]; ok && val != "" {
			value["create_time"] = time.Unix(val.(int64),0).Format("2006-01-02 15:04:05")
		}
	}
	return user,nil,count
}

func DelUserById(id int)(int,error)  {
	result,err := bootstrap.GetDb().Table("user").Where(map[string]interface{}{"id":id}).Delete()
	if err != nil {
		return  0,err
	}
	return  result,nil
}

func EditUserStatus(id int,status int)(bool,error)  {
	updateData := map[string]interface{}{
		"status" : status,
	}
	result,err := bootstrap.GetDb().Table("user").Data(updateData).
		Where(map[string]interface{}{"id":id}).Update()
	if err != nil {
		log.Fatal(err)
		return false,err
	}
	if result >= 1{
		return true,nil
	}
	return  false,nil
}

func SearchUserCondition(search service.Search)([][]interface{})  {
	var where  =  [][]interface{}{{"1","=","1"}}
	var symbol = map[int]string{0:">=",1:"<="}
	strType := reflect.TypeOf(search.Condition)
	strValue:= reflect.ValueOf(search.Condition)
	for k := 0;k < strType.NumField();k++{
		key := strings.ToLower(strType.Field(k).Name)
		value:= strValue.Field(k).Interface()
		if res,ok := value.([]int);ok{
			for k,v := range res{
				where = append(where,[]interface{}{key,symbol[k],v / 1000})
			}
		}else if value != ""{
			where = append(where,[]interface{}{key,"=",value})
		}
	}
	return  where
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