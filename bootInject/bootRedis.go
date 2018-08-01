package bootInject

import (
	"go-blog/bootstrap"
	"github.com/garyburd/redigo/redis"
)

func BootRedis() func(boot *bootstrap.Boot) {
	return func(boot *bootstrap.Boot) {
		boot.Redis,_ =  redis.Dial("tcp", "127.0.0.1:6379")
	}
}
