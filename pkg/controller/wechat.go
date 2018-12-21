package controller

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"html"
	"io/ioutil"
	"sort"
	"strings"
	"wechat/pkg/logs"
	"wechat/pkg/service"
	"wechat/pkg/service/message"
	"wechat/pkg/tool"
)

const (
	aesKey = "MhunqKssXObQ53QhlNPIIyhQftzoI3VOXoOwMRWyiYU";
	token  = "Macarinal"
)

type WeChat struct {
	params []byte
}

func NewWeChat() *WeChat {
	instance := new(WeChat)
	return instance
}

// 验证
func (this *WeChat)EntryPoint(ctx *gin.Context) {
	params    := make([]string, 0)
	signature := html.EscapeString(ctx.Query("signature"))
	timestamp := html.EscapeString(ctx.Query("timestamp"))
	nonce 	  := html.EscapeString(ctx.Query("nonce"))
	echoStr   := html.EscapeString(ctx.Query("echostr"))
	params = append(append(append(params, token), timestamp), nonce)
	sort.Strings(params)
	result := strings.Join(params, "")
	result = tool.Sha1Encrypt(result)
	if signature == result {
		if echoStr == ""{
			resp,err := ioutil.ReadAll(ctx.Request.Body)
			this.params = resp
			reply,err := this.MessageCenter()
			if err != nil {
				logs.Zap.Error(err.Error())
			}
			ctx.String(200,reply)
		}else{
			ctx.String(200,echoStr)
		}
	}
}

// 消息处理
func (this *WeChat) MessageCenter()(string,error)  {
	// 解析xml
	msg := message.Message{}
	err := xml.Unmarshal(this.params,&msg)
	if err != nil{
		return "",err
	}
	return service.NewWeChatService().MessageHandler(msg)
}
