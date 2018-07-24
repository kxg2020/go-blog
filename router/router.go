package router

import (
	"github.com/gin-gonic/gin"
	"go-blog/controller/admin"
	"strings"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"go-blog/controller/index"
)

func RouterInit() *gin.Engine {
	router := gin.Default()
	router.Use(CrossSite())

	router.Static("/static","./static")
	adminRouter := router.Group("/admin")
	adminRouter.POST("/article/upload",admin.NewArticle().Upload)
	adminRouter.POST("/login/validate",admin.NewLogin().Validate)
	adminRouter.Use(JwtAuth())
	adminRouter.POST("/article/cover",admin.NewArticle().UploadCover)
	adminRouter.POST("/check",admin.NewIndex().CheckToken)
	adminRouter.GET("/index",admin.NewIndex().Index)
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

	indexRouter := router.Group("/index")
	indexRouter.GET("/tag/list",index.NewTag().GetTagList)
	indexRouter.POST("/article/list",index.NewArticle().GetArticleList)
	indexRouter.GET("/article/info/:id",index.NewArticle().GetArticleInfo)

	return  router
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
		headerStr += ",Authorization"
		if headerStr != "" {
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Headers," + headerStr
		}else {
			headerStr = "Access-Control-Allow-Origin, Access-Control-Allow-Headers"
		}
		if origin == "http://127.0.0.1:8080" || origin == "http://127.0.0.1:8081" {
			// 定义可以被暴露给客户端响应头
			expose := "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers"
			expose += ",Content-Type,user-login-expire,user-login-token";
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Max-Age", "600")

			c.Header("Access-Control-Expose-Headers", expose)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Cache-Control", "no-store")
			c.Set("Content-Type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(204, []byte(""))
		}
		c.Next()
	}
}
