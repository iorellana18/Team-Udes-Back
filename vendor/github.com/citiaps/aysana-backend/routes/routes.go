package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	palabras := app.Group("/palabras")
	{
		palabras.GET("/categoria", PalabrasPorCategoria)
	}
	
	localidad := app.Group("/localidad")
	{
		localidad.GET("/region", RegionPorCategoria)
		localidad.GET("/comuna", ComunaPorCategoria)
	}
	
	tweet := app.Group("/tweet")
	{
		tweet.GET("/categoria", TweetsPorCategoria)
		tweet.GET("/categoria/count", CantTweetsPorCategorias)
	}
}