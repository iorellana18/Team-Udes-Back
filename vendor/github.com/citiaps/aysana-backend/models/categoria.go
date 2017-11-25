package models

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"net/http"
	"strconv"

	"github.com/citiaps/aysana-backend/db"
	"github.com/citiaps/aysana-backend/utils"
)

type Categoria struct {
	//	gorm.Model
	Id     int
	Nombre string
}

func CategoriaCRUD(app *gin.Engine) {
	app.GET("/categoria/:id", CategoriaFetchOne)
	app.GET("/categoria/", CategoriaFetchAll)
	app.POST("/categoria/", CategoriaCreate)
	app.PUT("/categoria/:id", CategoriaUpdate)
	app.DELETE("/categoria/:id", CategoriaRemove)
}

func CategoriaFetchOne(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	utils.Check(err)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("categoria", database)

	var categoria Categoria
	if err := collection.Find(bson.M{}).Skip(num).One(&categoria); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, categoria)
	}
}

func CategoriaFetchAll(c *gin.Context) {
	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("categoria", database)

	var categoria []Categoria
	if err := collection.Find(bson.M{}).All(&categoria); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, categoria)
	}
}

func CategoriaCreate(c *gin.Context) {
	var categoria Categoria
	err := c.BindJSON(&categoria)
	utils.Check(err)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("categoria", database)

	if err := collection.Insert(&categoria); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, categoria)
	}
}

func CategoriaUpdate(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}

func CategoriaRemove(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}
