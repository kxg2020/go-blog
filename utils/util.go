package utils

import (
	"time"
	"math/rand"
	"os"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RandInt64(min,max int64) int64{
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 判断是否是ajax
func IsAjax(ctx *gin.Context)(bool) {
	if ctx.GetHeader("X-Requested-With") != ""{
		return true
	}
	return  false
}

// 类型转换
func ToString(value interface{}) string {
	var result string
	switch value.(type) {
	case string:
		result = value.(string)
	case int:
		result = strconv.Itoa(value.(int))
	}
	return result
}
