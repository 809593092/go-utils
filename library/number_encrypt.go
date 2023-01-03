package library

import "github.com/speps/go-hashids/v2"

// NumberEncrypt 加密
// @demo NumberEncrypt("123", 18, []int{1001})
func NumberEncrypt(salt string, minLength int, params []int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err == nil {
		e, err := h.Encode(params)
		if err == nil {
			return e
		}
	}
	return ""
}

// NumberDecrypt 解密
// @demo NumberDecrypt("123", 18, "00898fddf8d9f8ds9f89")
func NumberDecrypt(salt string, minLength int, hash string) []int {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err == nil {
		e, err := h.DecodeWithError(hash)
		if err == nil {
			return e
		}
	}
	return []int{}
}
