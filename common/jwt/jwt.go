package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

//签名用的secret
var mySecret = []byte("cjq")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 自定义结构体,添加自定义的字段
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "cjq",
		},
	}
	// 使用指定的签名方法创建签名对象(加密算法,token配置)
	//header payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 签名并获得完整的编码后的字符串token
	//signature
	return token.SignedString(mySecret)
	//最终的base64编码
}

// 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token是否有效
		return mc, nil
	}
	return nil, errors.New("无效的token")
}
