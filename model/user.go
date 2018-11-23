package model

import (
	"meeting/bootstrap"
	"strings"
)

type User struct {
	Id          int64  `orm:"id"           json:"id,omitempty" `
	User_id     string `orm:"user_id"      json:"user_id,omitempty"`
	Name        string `orm:"name"         json:"name,omitempty"`
	Department  string `orm:"department"   json:"department,omitempty"`
	Avatar      string `orm:"avatar"       json:"avatar,omitempty"`
	Mobile      string `orm:"mobile"       json:"mobile,omitempty"`
	Position    string `orm:"position"     json:"position,omitempty"`
}

// 构造函数
func NewUser()*User  {
	return new(User)
}

func (this *User)UserList(pgNum int,pgSize int,user User)([]User,int64,error) {
	var users []User
	start := (pgNum - 1) * pgSize
	where := this.parseWhere(user)
	err   := bootstrap.Rose().Table(&users).Where(where).Offset(start).Limit(pgSize).Select()
	count,err := bootstrap.Rose().Table("user").Where(where).Count("id")
	if err != nil{
		return users,count,err
	}

	return users,count,nil
}

func (this *User)UserUpdate(params User)(int,error)  {
	go func(params User) {

	}(params)
	return  1,nil
}

// where条件
func (this *User) parseWhere(params User)([][]interface{})  {
	where := [][]interface{}{}
	where = append(where,[]interface{}{"1","=","1"})
	if params.Name != "" {
	 	where = append(where, []interface{}{"name","like","%"+strings.TrimSpace(params.Name)+"%"})
	}
	if params.Mobile != "" {
		where = append(where,[]interface{}{"mobile","=",strings.TrimSpace(params.Mobile)})
	}
	if params.Position != "" {
		where = append(where,[]interface{}{"position","=",strings.TrimSpace(params.Position)})
	}
	return where
}

