package robot

import (
	"encoding/json"
	"wechat/pkg/tool"
)

const ApiKey    = "709a5f098e7545e09c420b410feb9deb"
const Api       = "http://openapi.tuling123.com/openapi/api/v2"

type perception struct {
	InputText  map[string]string   `json:"inputText"`
	InputImage map[string]string   `json:"inputImage"`
	SelfInfo   SelfInfo 		   `json:"selfInfo"`
}

type Location struct {
	City 	  string     `json:"city"`
	Province  string     `json:"province"`
	Street    string     `json:"street"`
} 

type SelfInfo struct {
	Location Location     `json:"location"`
}

type userInfo struct {
	ApiKey    string      `json:"apiKey"`
	UserId    string      `json:"userId"`
}

type TuLing struct {
	ReqType    int 		  `json:"reqType"`
	Perception perception `json:"perception"`
	UserInfo   userInfo   `json:"userInfo"`
}

func NewTuLing(params map[string]string) *TuLing  {
	robot := &TuLing{
		ReqType : 0,
		Perception:perception{
			map[string]string{"text" : params["msg"]},
			map[string]string{"url":""},
			SelfInfo{
			Location{"四川", "成都", "",
					},
			},
		},
		UserInfo:userInfo{ApiKey,"macarinal"},
	}
	return robot
}

func (this *TuLing)RobotAnswer() (map[string]interface{},error) {
	var result map[string]interface{}
	postStr,err := json.Marshal(this)
	if err != nil {
		return result,err
	}
	ret := tool.HttpPostJson(Api,postStr)
	err  = json.Unmarshal(ret,&result)
	if err != nil {
		return result,err
	}
	return result,nil
}
