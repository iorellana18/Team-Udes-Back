package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/citiaps/aysana-backend/db"
	"github.com/citiaps/aysana-backend/models"
	"github.com/citiaps/aysana-backend/routes"
	"github.com/citiaps/aysana-backend/utils"
)

func main() {
	utils.LoadConfig("config/config.yaml")

	db.ElasticSetup()
	db.MongoSetup()

	app := gin.Default()
	app.Use(cors.Default())

	routes.Setup(app)
	models.Setup(app)

	app.Run(":" + utils.Config.Server.Port)
}
