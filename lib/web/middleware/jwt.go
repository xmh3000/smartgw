package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var jwtSecret = []byte("smartgw.com.cn")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func GenerateToken(username string) (string, error) {
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			Issuer:    "smart.com.cn",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("解析错误！")
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")

		if token == "" {
			context.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "没有携带token",
				"data":    "",
			})
			context.Abort()
			return
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "token验证失败",
					"data":    "",
				})
				context.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				context.JSON(http.StatusOK, gin.H{
					"code":    http.StatusUnauthorized,
					"message": "token已过期",
					"data":    "",
				})
				context.Abort()
				return
			}
		}
	}
}
