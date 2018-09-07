package bootstrap
import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
	_"github.com/go-sql-driver/mysql"
	"log"
)
type Bootstrap struct {
	Db     *gorose.Database
	Router *gin.Engine
}
var Boot *Bootstrap
var config = map[string]interface{} {
	"Default"         : "mysql_dev",
	"SetMaxOpenConns" : 20,
	"SetMaxIdleConns" : 2,
	"Connections"     : map[string]map[string]string{
		   "mysql_dev": {
			"host"    : "127.0.0.1",
			"username": "root",
			"password": "root",
			"port"    : "3306",
			"database": "xm_blog",
			"charset" : "utf8",
			"protocol": "tcp",
			"prefix"  : "xm_",
			"driver"  : "mysql",
		},
	},
}

// 初始化数据库
func (this *Bootstrap)initDb() *Bootstrap{
	var err error
	db,err := gorose.Open(config)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	this.Db = db.GetInstance()
	return this
}

// 初始化框架
func (this *Bootstrap)initFramework() *Bootstrap {
	this.Router = gin.Default();
	return this
}

// 启动项目
func InitBootstrap() *Bootstrap{
	return new(Bootstrap).initDb().initFramework().setBootstrap()
}

// 保存对象
func (this *Bootstrap)setBootstrap() *Bootstrap {
	Boot = this
	return this
}

// 获取数据库连接
func GetDb() *gorose.Database {
	return Boot.Db
}