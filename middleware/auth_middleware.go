package middleware

import (
	"net/http"
	"strings"
	"wms-prototype/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Masuk Middleware")
		//ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		println("AUTH HEADER:", authHeader)

		//cek apakah header kosong

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}
		//format yang diharapkan :
		//bearer eyJhbGciOiJIUzI1NiIs
		tokenstring := strings.TrimPrefix(authHeader, "Bearer ")

		//parse dan validasi token
		token, err := jwt.Parse(
			tokenstring,
			func(t *jwt.Token) (interface{}, error) {

				//pastikan algoritma yang digunakan adalah HS256
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return utils.Secretkeys, nil
			},
		)

		//cek token
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}
		//ambil claims

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		//simpan ke context
		c.Set("userID", uint(claims["user_id"].(float64)))
		c.Set("email", claims["email"].(string))
		c.Set("role", claims["role"].(string))

		c.Next()

	}
}
