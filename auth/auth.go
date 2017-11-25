package auth

import (
	"github.com/gin-gonic/gin"
	jwt "gopkg.in/appleboy/gin-jwt.v2"

	"time"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/models"
)

func Authenticator(userId string, password string, c *gin.Context) (string, bool) {
	db := db.Database()
	defer db.Close()

	var user models.Users

	if err := db.Where("email = ?", userId).First(&user).Error; err != nil {
		return userId, false
	} else {
		if password == user.Password {
			return userId, true
		} else {
			return userId, false
		}
	}
}

func Authorizator(userId string, c *gin.Context) bool {
	typeUser := typeUser(userId)
	if authGroups, exists := c.Get("authGroups"); exists {
		if ok := authorizatorUser(typeUser, authGroups.([]int)); ok {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func CreateMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "Users",
		Key:           []byte("ehackathon"),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		Authenticator: Authenticator,
		Authorizator:  Authorizator,
		Unauthorized:  Unauthorized,

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	return authMiddleware
}
