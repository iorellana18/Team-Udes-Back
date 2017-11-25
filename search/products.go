package search

import (
	"github.com/gin-gonic/gin"

	elastic "gopkg.in/olivere/elastic.v5"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/models"
	"net/http"
	"reflect"
	"strconv"
)

func QueryIndex(c *gin.Context, query string, paginacion int) ([]models.ProductIndex, error) {
	ctx, client := db.ElasticInit()

	categoriaQuery := elastic.NewTermQuery("title", query)
	if searchResult, err := client.Search().Index(db.GetIndex()).Query(categoriaQuery).From(10 * paginacion).Size(10).Do(ctx); err != nil {
		return nil, err
	} else {
		products := make([]models.ProductIndex, len(searchResult.Hits.Hits))
		for i, item := range searchResult.Each(reflect.TypeOf(models.ProductIndex{})) {
			productsRes := item.(models.ProductIndex)
			products[i] = productsRes
		}

		if len(searchResult.Hits.Hits) > 0 {
			return products, nil
		} else {
			return nil, nil
		}
	}
}

func SearchProduct(c *gin.Context) {
	query := c.Query("q")

	scroll := c.DefaultQuery("scroll", "0")
	if paginacion, err := strconv.Atoi(scroll); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		if products, err := QueryIndex(c, query, paginacion); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			if products != nil {
				c.JSON(http.StatusOK, products)
			} else {
				c.String(http.StatusNotFound, "Not Found")
			}
		}
	}
}
