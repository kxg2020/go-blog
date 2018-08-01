package bootInject

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/gzip"
	"go-blog/controller/admin"
	"go-blog/controller/index"
	"strings"
	"go-blog/bootstrap"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func BootGin() func(boot *bootstrap.Boot)  {
	return func(boot *bootstrap.Boot) {
		boot.Router = gin.Default()
		boot.Router.Use(CrossSite())
		boot.Router.Use(gzip.Gzip(gzip.DefaultCompression))
		adminRouter := boot.Router.Group("/admin")
		adminRouter.POST("/login/validate",admin.NewLogin().Validate)
		adminRouter.Use(JwtAuth())
		{
			adminRouter.POST("/article/cover",admin.NewArticle().UploadCover)
			adminRouter.POST("/check",admin.NewIndex().CheckToken)
			adminRouter.Any("/article/insert",admin.NewArticle().Insert)
			adminRouter.Any("/article/list",admin.NewArticle().List)
			adminRouter.GET("/article/info/:id",admin.NewArticle().ArticleInfo)
			adminRouter.Any("/article/editStatus",admin.NewArticle().EditStatus)
			adminRouter.GET("/article/delete/:id",admin.NewArticle().Delete)
			adminRouter.POST("/article/saveEdit",admin.NewArticle().SaveEdit)
			adminRouter.GET("/user/list",admin.NewUser().GetUserList)
			adminRouter.POST("/user/edit",admin.NewUser().EditUserStatus)
			adminRouter.POST("/user/saveEdit",admin.NewUser().SaveUseEdit)
			adminRouter.POST("/user/delete",admin.NewUser().DelUser)
			adminRouter.POST("/user/insert",admin.NewUser().Insert)
			adminRouter.GET("/tag/list",admin.NewTag().GetTagList)
			adminRouter.POST("/tag/insert",admin.NewTag().InsertTag)
			adminRouter.Any("/tag/editStatus/:id",admin.NewTag().EditTagStatus)
			adminRouter.POST("/tag/editTag",admin.NewTag().EditTag)
			adminRouter.POST("/tag/delete",admin.NewTag().DelTag)
		}


		indexRouter := boot.Router.Group("/index")
		indexRouter.GET("/tag/list",index.NewTag().GetTagList)
		indexRouter.POST("/article/list",index.NewArticle().GetArticleList)
		indexRouter.GET("/article/info/:id",index.NewArticle().GetArticleInfo)
	}
}

// jwt验证中间件(jwt只验证token是否存在且是否正确,是就认为是真实的登陆状态,并执行到下一步next,即跳转到用户请求的那个控制器)
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("token"), nil
			})
		if err == nil {

			if token.Valid {
				c.Next()
			}else {
				c.AbortWithStatusJSON(200,map[string]interface{}{
					"status" : 2,
					"msg"    : "token已经失效,请重新登陆",
				})
			}
		}else{
			c.AbortWithStatusJSON(200,map[string]interface{}{
				"status" : 2,
				"msg"    : "token非法,请重新登陆",
			})
		}
	}
}

// 跨域
func CrossSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		headerStr += ",Authorization,Content-Type"
		if headerStr != "" {
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Headers," + headerStr
		}else {
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Headers"
		}
		fmt.Println(headerStr)
		if origin == "http://127.0.0.1:8080" || origin == "http://127.0.0.1:8081" {
			// 暴露给客户端的响应头
			expose := "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers"
			expose += ",Content-Type,user-login-expire,user-login-token";
			// 允许请求的域
			c.Header("Access-Control-Allow-Origin", origin)
			// 允许请求的header头
			c.Header("Access-Control-Allow-Headers", headerStr)
			// 允许请求的方法类型
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// 上面设置的允许内容的缓存时间,在时间范围内,浏览器不会发起第二次options请求
			c.Header("Access-Control-Max-Age", "1800")
			// 返回给客户端的header头
			c.Header("Access-Control-Expose-Headers", expose)
			// 是否验证cookie
			c.Header("Access-Control-Allow-Credentials", "true")
			// 返回的数据内容是否缓存
			c.Header("Cache-Control", "no-store")
			// 返回的数据格式
			c.Set("Content-Type", "application/json")
		}

		// 不处理OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
		}else{
			c.Next()
		}

	}
}