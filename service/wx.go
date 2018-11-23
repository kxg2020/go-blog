package service

import (
	"encoding/json"
	"fmt"
	"log"
	"meeting/util"
)

// 应用ID
const AGENT_ID     = 1000003;
// 应用秘钥
const AGENT_SECRET = "RMJaW3iZxl7T0plKk4qG_qaaAIkM4xeWz2S5HmfNzxk";
// 企业ID
const COMPANY_ID   = "ww0d4cee19e94134fd";
// 企业接口
const COMPANY_BASE_API  = "https://qyapi.weixin.qq.com/cgi-bin/";
// 获取access_token
const GET_ACCESS_TOKEN  = "gettoken?corpid=%s&corpsecret=%s";
// 获取成员基础信息
const GET_MEMBER_BASIC  = "user/getuserinfo?access_token=%s&code=%s";
// 获取成员详细信息
const GET_MEMBER_INFO   = "user/get?access_token=%s&userid=%s";
// 获取部门列表信息
const GET_DEPARTMENT    = "department/list?access_token=%s&id=%s";
// 发送应用卡片消息
const SEND_AGENT_MESSAGE= "message/send?access_token=%s";
// 获取部门成员
const GET_DEPARTMENT_MEMBER = "user/simplelist?access_token=%s&department_id=%s&fetch_child=%s";

// departmentList
type Wx struct {
	Url    string
	Token  string
}

// 构造
func NewWx()*Wx  {
	return new(Wx)
}

// token
func (this *Wx)GetToken() *Wx {
	var token Token
	this.requestRouter("token").request("GET",&token,map[string]interface{}{})
	this.Token = token.Access_token
	return this
}

// department
func (this *Wx)GetDepartmentList()DepartmentList  {
	var department DepartmentList
	this.GetToken().requestRouter("departmentList").request("GET",&department,map[string]interface{}{})
	return department
}

func(this *Wx) requestRouter(route string)*Wx  {
	switch route{
	case "token":
		this.Url = fmt.Sprintf(COMPANY_BASE_API+GET_ACCESS_TOKEN,COMPANY_ID,AGENT_SECRET)
	case "departmentList":
		this.Url = fmt.Sprintf(COMPANY_BASE_API+GET_DEPARTMENT,this.Token,"")
	}
	return this
}

func (this *Wx)request(method string,container interface{},data map[string]interface{})  {
	result,err := util.HttpRequest(method,this.Url,data)
	if err != nil{
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(result,container)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
}

