package admin

import (
	"github.com/gin-gonic/gin"
	"go-blog/utils"
	"go-blog/model"
	"log"
	"strconv"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Login struct {

}
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}


// 构造函数
func NewLogin() *Login {
	login := new(Login)
	return login
}

// 登陆首页
func (login *Login)Index(ctx *gin.Context)  {
	ctx.HTML(200,"admin/login/index.html",map[string]string{})
}

// 验证登陆
func (login *Login)Validate(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	userInfo,err := model.GetUserInfo(username)
	if err != nil{
		log.Fatal(err.Error())
		utils.PrintErrors(4000,ctx)
		return
	}

	if len(userInfo) > 0{
		// 密码是否相等
		passwordMd5 := utils.NewEncrypt().Md5(password + userInfo["salt"].(string))
		if passwordMd5 != userInfo["password"]{
			utils.PrintErrors(4000,ctx)
			return
		}
		// 更新盐和密码
		saltNew     := utils.RandInt64(0,999)
		passwordNew := utils.NewEncrypt().Md5(password + strconv.FormatInt(saltNew,10))
		updateRes   := model.UpdateUserPassword(saltNew,passwordNew,userInfo["id"].(int64))
		if updateRes{
			ctx.Header("user-login-token",setToken())
			ctx.Header("user-login-expire","2")
			utils.PrintSuccess(9000,map[string]string{
				"username":userInfo["username"].(string),
			},ctx)
			return
		}
	}else{
		utils.PrintErrors(4000,ctx)
	}
}

// 设置jwtToken
func setToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(2)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("token"))
	if err != nil {
		return ""
	}
	return tokenString
}
