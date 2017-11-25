package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iorellana18/Team-Udes-Back/auth"
)

func Setup(app *gin.Engine) {
	authMiddleware := auth.CreateMiddleware()

	app.POST("/login", authMiddleware.LoginHandler)

	authorization := app.Group("/auth")
	authorization.Use(auth.AddPermission(auth.Admin, auth.User))
	authorization.Use(authMiddleware.MiddlewareFunc())
	{
		authorization.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
}
