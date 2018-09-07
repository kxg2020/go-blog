package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密
func Md5Encrypt(str string)string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return  hex.EncodeToString(ctx.Sum(nil))
}
