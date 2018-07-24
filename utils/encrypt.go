package utils

import (
	"crypto/md5"
	"encoding/hex"
)

type Encrypt struct {

}

// 构造函数
func NewEncrypt() *Encrypt {
	encrypt := new(Encrypt)
	return  encrypt
}

// md5加密
func (encrypt *Encrypt)Md5(param string) string{
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(param))
	result := md5Ctx.Sum(nil)
	return hex.EncodeToString(result)
}


