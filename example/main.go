package main

import (
	"github.com/809593092/go-utils/hash"
	"github.com/809593092/go-utils/library"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	// go utils使用方法
	// md5封装的方法
	println(hash.StringMd5("123456"))

	// token 加密
	key := hash.StringMd5("123456")
	library.EncodeToken(key, jwt.MapClaims{"username": "mark", "ip": "127.0.0.1"})
}
