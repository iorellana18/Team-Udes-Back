package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iorellana18/Team-Udes-Back/api"
	"github.com/iorellana18/Team-Udes-Back/auth"
	"github.com/iorellana18/Team-Udes-Back/search"
)

func Setup(app *gin.Engine) {
	authMiddleware := auth.CreateMiddleware()

	app.POST("/login/", authMiddleware.LoginHandler)

	authorization := app.Group("/auth")
	authorization.Use(auth.AddPermission(auth.Admin, auth.User))
	authorization.Use(authMiddleware.MiddlewareFunc())
	{
		authorization.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	img := app.Group("/img")
	{
		img.POST("/analyze/", api.AnalyzeImagen)
	}

	app.GET("/search", search.SearchProduct)
}
