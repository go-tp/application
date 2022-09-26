package extend

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "gtp"
)

/*
传入的是用户数据
用户数据
claims := jwt.MapClaims{
	"Username":   1,
	"Password": "test",
	"StandardClaims" : jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
		ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
		Issuer:    "newtrekWang",                   //签名的发行者
	}
}
*/
// 创建token
func CreateJwt(claims jwt.MapClaims) (string, error) {

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(SECRETKEY))

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

// 解析token
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}