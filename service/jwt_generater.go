package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"message-board/utils"
	"time"
)

var jwtSecret = utils.JwtSecret // 用于签名和验证 Token 的密钥

func GenerateJWT(userID int) (string, error) {
	// 创建 JWT
	fmt.Println(userID)
	fmt.Println(string(jwtSecret))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                                // 获取用户id
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // 设置 Token 过期时间
	})

	// 使用密钥签名 Token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
