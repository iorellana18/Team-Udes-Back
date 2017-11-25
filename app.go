package main

import (
	"github.com/gin-gonic/gin"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/models"
	"github.com/iorellana18/Team-Udes-Back/routes"
	"github.com/iorellana18/Team-Udes-Back/utils"
	//"github.com/iorellana18/Team-Udes-Back/migrate"
)

func main() {
	utils.LoadConfig("config/config.yaml")

	db.MongoSetup()
	db.ElasticSetup()
	db.PostgresSetup()

	app := gin.Default()
	app.Use(utils.CorsMiddleware())

	models.Setup(app)
	routes.Setup(app)

	//migrate.MongoToElastic();

	app.Run(":" + utils.Config.Server.Port)
}
