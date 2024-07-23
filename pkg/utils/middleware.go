package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aldyN25/myproject/app/config"
	"github.com/aldyN25/myproject/app/models"
	gormdb "github.com/aldyN25/myproject/pkg/gorm.db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ExtractClaims(secret, tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte(secret)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid JWT Token")
}

func CheckAuth(c *gin.Context) {
	configs := config.GetInstance()
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.Jwtconfig.Secret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	db, err := gormdb.GetInstance()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var users *models.Users
	err = db.Debug().Where("user_id = ?", claims["userID"]).Find(&users).Error
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if users == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("userID", claims["userID"])

	c.Next()

}
