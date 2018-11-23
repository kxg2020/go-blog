package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"go/src/github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// 分页格式
func PageFormat(pgNum,pgSize int)(int,int) {

	if pgNum > 100 || pgNum < 1 {
		pgNum = 1
	}
	if pgSize >15 || pgSize < 1 {
		pgSize = 9
	}
	return pgNum,pgSize
}

// jsonDecode
func JsonDecode(params string,args interface{}) (interface{},error) {
	err := json.Unmarshal([]byte(params),&args)
	if err != nil{
		return nil,nil
	}
	return args,nil
}

// 长度检测
func CheckLength(params string)bool  {
	if len(params) == 0{
		return  false
	}
	return true
}

// MD5加密
func Md5Encrypt(params string)string  {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(params))
	md5String := md5Ctx.Sum(nil)
	return hex.EncodeToString(md5String)
}

// 随机数
func RandNumber(min int64,max int64)int64  {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max - min)
}

// Token
func CreateToken(key string,username string)string  {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := make(jwt.MapClaims)
	claim["exp"] = time.Now().Add(time.Duration(24) * time.Hour).Unix()
	claim["iat"] = time.Now().Unix()
	claim["username"] = username
	token.Claims = claim
	tokenString,err := token.SignedString([]byte(key))
	if err != nil{
		log.Fatal(err.Error());
		return ""
	}
	return tokenString
}

// Token Validate
func TokenValidate() gin.HandlerFunc  {
	return func(context *gin.Context) {
		token,err := request.ParseFromRequest(context.Request,request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte("meeting"),nil
		})
		if err != nil{
			context.AbortWithStatusJSON(200,map[string]interface{}{
				"status" : 2,
				"msg"    : "token非法,请重新登陆",
			})
		}else{
			if token.Valid{
				if claim,ok := token.Claims.(jwt.Claims);ok{
					context.Set("username",claim.(jwt.MapClaims)["username"])
				}
				context.Next()
			}else{
				context.AbortWithStatusJSON(200,map[string]interface{}{
					"status" : 2,
					"msg"    : "token已经失效,请重新登陆",
				})
			}
		}
	}
}


func HttpRequest(method string,url string,data map[string]interface{})([]byte,error)  {
	var req   http.Request
	var resp *http.Response
	req.ParseForm()
	if len(data) != 0 {
		for key,val := range data{
			req.Form.Add(key, val.(string))
		}
	}
	bodyStr := strings.TrimSpace(req.Form.Encode())
	main, err := http.NewRequest(method, url, strings.NewReader(bodyStr))
	if err != nil {
		return nil,err
	}
	main.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	main.Header.Set("Connection", "Keep-Alive")
	resp, err = http.DefaultClient.Do(main)
	if err != nil {
		return nil,err
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	return body,nil
}

