package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"time"
	"backendApi/model"
)

// 跨域
func CrossSite() gin.HandlerFunc {
	return func(context *gin.Context) {
		var headers []string
		var headerStr string
		var expose    string
		method := context.Request.Method
		origin := context.Request.Header.Get("origin")
		for key,_ := range context.Request.Header{
			headers = append(headers,key)
		}
		// 数组转字符串
		headerStr = strings.Join(headers,",")
		headerStr+= ",Authorization,Content-Type,Username,UserOnline";
		if headerStr != ""{
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Headers," + headerStr
		}else{
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Headers";
		}

		if origin == "http://127.0.0.1:8080" {
			// 暴露给客户端的响应头
			expose  = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers"
			expose += ",Content-Type,user-expire,user-token,user-online";
			// 允许请求的域
			context.Header("Access-Control-Allow-Origin", origin)
			// 允许请求的header头
			context.Header("Access-Control-Allow-Headers", headerStr)
			// 允许请求的方法类型
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// 上面设置的允许内容的缓存时间,在时间范围内,浏览器不会发起第二次options请求
			context.Header("Access-Control-Max-Age", "1800")
			// 返回给客户端的header头
			context.Header("Access-Control-Expose-Headers", expose)
			// 是否验证cookie
			context.Header("Access-Control-Allow-Credentials", "true")
			// 返回的数据内容是否缓存
			context.Header("Cache-Control", "no-store")
			// 返回的数据格式
			context.Set("Content-Type", "application/json")
		}

		if method == "OPTIONS"{
			context.AbortWithStatus(204)
		}else{
			context.Next()
		}
	}
}

const whiteMenu  = "/vi/login"

// jwt
func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Request.Header.Get("Username")
		userOnline := context.Request.Header.Get("UserOnline")
		if len(username) == 0 || username == ""{
			username = ""
		}
		token, err := request.ParseFromRequest(context.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("token"+username), nil
			})
		if err == nil {
			if token.Valid {
				if context.Request.URL.Path != whiteMenu{
					if LoginTokenValidate(username,userOnline){
						context.Next()
					}else{
						context.AbortWithStatusJSON(200,map[string]interface{}{
							"status" : 2,
							"code"   : 10000,
							"msg"    : "账号已在其他地方登陆",
						})
					}
				}
			}else {
				context.AbortWithStatusJSON(200,map[string]interface{}{
					"status" : 2,
					"code"   : 10000,
					"msg"    : "请重新登陆",
				})
			}
		}else{
			context.AbortWithStatusJSON(200,map[string]interface{}{
				"status" : 2,
				"code"   : 10000,
				"msg"    : "请重新登陆",
			})
		}
	}
}

// 设置jwt
func SetToken(username string)string  {
	token := jwt.New(jwt.SigningMethodHS256)
	claims:= make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(2)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims  = claims
	tokenString, err := token.SignedString([]byte("token"+username))
	if err != nil {
		return ""
	}
	return tokenString
}

func LoginTokenValidate(username string,token string)bool  {
	return  model.ValidateLoginToken(username,token)
}