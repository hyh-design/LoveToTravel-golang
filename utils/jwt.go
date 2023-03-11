package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// 想要保存更多信息都可以添加到这个结构体中
type MyClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(ID string, Name string) (string, error) {
	c := MyClaims{
		ID,
		Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt.expire")) * time.Hour).Unix(),
			Issuer: viper.GetString("auth.jwt.issuer"),
		},
	}
	// 使用指定的签名方法创建签名对象, 使用指定的secret签名并获得完整的编码后的字符串token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(viper.GetString("auth.jwt.secret")))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		fmt.Println("token parsing...")
		return []byte(viper.GetString("auth.jwt.secret")), nil
	})
	if err != nil {
		fmt.Println("token err!")
		fmt.Println(err)
		return nil, err
	}
	// 令牌有效
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		fmt.Println("token checked!")
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
