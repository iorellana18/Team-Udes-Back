package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/models"
)

func AddPermission(users ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("authGroups", users)
	}
}

func typeUser(userId string) int {
	db := db.Database()
	defer db.Close()

	var user models.Users

	if err := db.Where("email = ?", userId).First(&user).Error; err != nil {
		return 0
	} else {
		return user.User_type
	}
}

func authorizatorUser(typeUser int, authGroups []int) bool {
	for _, authGroup := range authGroups {
		if authGroup == typeUser {
			return true
		}
	}
	return false
}
