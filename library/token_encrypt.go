package library

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// const mySigningKey = "e10adc3949ba59abbe56e057f20f883e"

// EncodeToken 加密token
// @param key 加密的key
// @param mapClaims要加密的数据map
func EncodeToken(mySigningKey string, mapClaims jwt.MapClaims) (string, error) {
	// 创建
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	return claims.SignedString([]byte(mySigningKey))
}

// DecodeToken 解密Token字符串
func DecodeToken(mySigningKey string, tokenStr string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//
		//}
		return []byte(mySigningKey), nil
	})

	if err != nil {
		return nil, err
	}

	// if token.Valid {
	// 	fmt.Println("You look nice today")
	// } else if errors.Is(err, jwt.ErrTokenMalformed) {
	// 	fmt.Println("That's not even a token")
	// } else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
	// 	// Token is either expired or not active yet
	// 	fmt.Println("Timing is everything")
	// } else {
	// 	fmt.Println("Couldn't handle this token:", err)
	// }

	// 验证token
	if !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//u := *(*map[string]interface{})(unsafe.Pointer(&claims))
		return claims, nil
	}
	return nil, errors.New("token 异常")
}
