package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义一个名为jwtKey的字节数组，用于签名
var jwtKey = []byte("my_secret_key")

// 定义一个Claims结构体，包含了用户信息和JWT标准声明
type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	jwt.StandardClaims
}

// getToken函数用于生成JWT
func GetToken(username string, userId uint) string {
	// 创建一个新的JWT，使用HS256算法进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			// 设置令牌过期时间为24小时后
			ExpiresAt: time.Now().Add(time.Hour * 24 * 24).Unix(),
		},
	})
	// 对JWT进行签名，并将签名结果转换为字符串形式
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Error generating token string")
	}
	return tokenString
}

func ParseToken(tokenString string) (bool, string, uint) {

	// 使用claims解析token
	tkn, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("解析token时出错")
		return false, "", 0
	}
	if !tkn.Valid {
		fmt.Println("无效的token")
		return false, "", 0
	}

	// 提取claims
	claims, ok := tkn.Claims.(*Claims)
	if !ok {
		fmt.Println("无法解析claims")
		return false, "", 0
	}

	// 打印claims中的用户名
	fmt.Println(claims.Username)
	return true, claims.Username, claims.UserId
}
