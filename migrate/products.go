package migrate

import (
	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/models"

	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func getAllProducts() []models.Product {
	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("products", database)

	var products []models.Product
	if err := collection.Find(bson.M{}).All(&products); err != nil {
		return nil
	}

	return products
}

func indexProduct(product models.Product) {
	ctx, client := db.ElasticInit()

	productIndex := models.ProductIndex{
		ProductID:    product.ProductID,
		URL:          product.URL,
		Brand:        product.Brand,
		Title:        product.Title,
		Rating:       product.Rating,
		TotalReviews: product.TotalReviews,
		Published:    strconv.FormatBool(product.Published),
		Destacados:   product.Destacados,
		Colores:      product.Colores,
		Sizes:        product.Colores,
		Categorias:   product.Categorias,
		OnlineStock:  product.OnlineStock,
		Precio:       product.Precio,
	}

	db.ElasticIndex(ctx, client, productIndex)
}

func MongoToElastic() {
	products := getAllProducts()

	for _, product := range products {
		indexProduct(product)
	}
}
