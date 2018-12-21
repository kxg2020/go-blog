package service

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
	"wechat/pkg/service/message"
	"wechat/pkg/service/robot"
	"wechat/pkg/tool"
)

type WeChat struct {
	token   string
	message message.Message
}

const(
	AppId     = "wx1a2f9b4ac1a71a25"
	AppSecret = "4219bf126507a103fb8001db63fb1cf3"
)

const(
  Token = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)


func NewWeChatService()*WeChat  {
	return new(WeChat)
}

func (this *WeChat)MessageHandler(msg message.Message)(string,error) {
	this.message = msg
	switch msg.MsgType {
	case "video":

	case "image":

	case "text":
		return this.text()
	case "event":
	}
	return "",nil
}

func (this *WeChat)text() (string,error){
	robotAnswer,err :=  robot.NewTuLing(map[string]string{"msg":this.message.Content}).RobotAnswer();
	reply := message.Text{
		ToUserName  : this.message.FromUserName,
		FromUserName: this.message.ToUserName,
		CreateTime  : strconv.Itoa(int(time.Now().Unix())),
		MsgType     : this.message.MsgType,
		Content     : robotAnswer["results"].([]interface{})[0].(map[string]interface{})["values"].(map[string]interface{})["text"].(string),
	}
	result,err := xml.Marshal(reply)
	if err != nil {
		return "",err
	}
	return string(result),nil
}

func (this *WeChat)image()  {

}

func (this *WeChat)video()  {

}

func (this *WeChat)event()  {

}

func (this *WeChat)AccessToken()  {
	this.token = tool.HttpGet(fmt.Sprintf(Token,AppId,AppSecret))
	fmt.Println(this.token)
}
