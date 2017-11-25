package models

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"net/http"
	"strconv"

	"fmt"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/utils"
)

type ProductIndex struct {
	ProductID    string   `json:"productId"`
	URL          string   `json:"url"`
	Brand        string   `json:"brand"`
	Title        string   `json:"title"`
	Rating       int      `json:"rating"`
	TotalReviews int      `json:"totalReviews"`
	Published    string   `json:"published"`
	Destacados   []string `json:"destacados"`
	Colores      []string `json:"colores"`
	Sizes        []string `json:"sizes"`
	Categorias   []string `json:"categorias"`
	OnlineStock  int      `json:"onlineStock"`
	Precio       string   `json:"precio"`
}

type Product struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	ProductID    string        `json:"productId"`
	URL          string        `json:"url"`
	Brand        string        `json:"brand"`
	Title        string        `json:"title"`
	Rating       int           `json:"rating"`
	TotalReviews int           `json:"totalReviews"`
	Published    bool          `json:"published"`
	Destacados   []string      `json:"destacados"`
	Colores      []string      `json:"colores"`
	Sizes        []string      `json:"sizes"`
	Categorias   []string      `json:"categorias"`
	OnlineStock  int           `json:"onlineStock"`
	Precio       string        `json:"precio"`
}

func ProductCRUD(app *gin.Engine) {
	app.GET("/product/id/:id", ProductFetchID)
	app.GET("/product/skip/:id", ProductFetchOne)
	app.GET("/product/", ProductFetchAll)
	app.POST("/product/", ProductCreate)
	app.PUT("/product/:id", ProductUpdate)
	app.DELETE("/product/:id", ProductRemove)
}

func ProductFetchID(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.String(http.StatusNotFound, "Not ID Hex")
	}

	oid := bson.ObjectIdHex(id)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("products", database)

	var product Product
	if err := collection.FindId(oid).One(&product); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		fmt.Println(product.Id)
		c.JSON(http.StatusOK, product)
	}
}

func ProductFetchOne(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	utils.Check(err)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("products", database)

	var product Product
	if err := collection.Find(bson.M{}).Skip(num).One(&product); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func ProductFetchAll(c *gin.Context) {
	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("products", database)

	var product []Product
	if err := collection.Find(bson.M{}).All(&product); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func ProductCreate(c *gin.Context) {
	var product Product
	err := c.BindJSON(&product)
	utils.Check(err)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("products", database)

	if err := collection.Insert(&product); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func ProductUpdate(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}

func ProductRemove(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}
