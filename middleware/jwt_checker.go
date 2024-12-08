package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"message-board/utils"
	"net/http"
)

var jwtSecret = utils.JwtSecret

// JWTAuthMiddleware 中间件
func JWTAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头中获取 Authorization 字段
		tokenString := c.GetHeader("Authorization")
		fmt.Println(string(tokenString))
		if string(tokenString) == "" { //没有token
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": utils.MissingToken.Error(),
			})
			c.Abort() // 中断后续流程
			return
		}

		// 解析并验证 Token
		token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
			// 确保签名方法是我们支持的 HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError(utils.InvalidTokenSingingMethod.Error(), jwt.ValidationErrorSignatureInvalid)
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": utils.InvalidToken.Error(),
			})
			c.Abort() // 中断后续流程
			return
		}

		// 将解析出的用户信息存入上下文，供后续使用
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			fmt.Printf("%T", claims["user_id"])
			fmt.Println("Claims:", claims) // 打印出所有解析到的 claims
		} else {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"error": utils.InvalidClaims.Error(),
			})
			c.Abort()
		}
	}
}
