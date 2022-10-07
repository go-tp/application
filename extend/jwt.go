package extend

import (
	"github.com/dgrijalva/jwt-go"
	"gtp/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
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


// 生成jwt
func GenerateToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user,
    // "exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})
	secret := configs.ReadYaml().JWT.Secret
	return token.SignedString([]byte(secret))//对应的字符串请自行生成，最后足够使用加密后的字符串
}

// 中间件身份认证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
	   //拿到token
	   tokenStr := c.Request.Header.Get("Authorization")
	   
	   if tokenStr == "" {
		  c.JSON(http.StatusOK, gin.H{
			 "status": -1,
			 "msg":    "token为空，请携带token",
			 "data":   nil,
		  })
		  c.Abort()
		  return
	   }
	   tokenStr = tokenStr[len("Bearer "):]

	   

	   token1,_ := ParseToken(tokenStr)
	   username,err := token1["username"]
	   if(err != true){
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "token错误",
				"data":   nil,
			})
			c.Abort()
			return
	   }
	   // 中间件传参 使用c.Get("username") 获取
	   c.Set("username",username)
	//    fmt.Println("username:",username)
	//    fmt.Println("err:",err)
		return
	}
 }


 // jwt解密
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	secret := configs.ReadYaml().JWT.Secret
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}