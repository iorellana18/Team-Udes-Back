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

func SearchProduct(c *gin.Context) {
	query := c.Query("q")

	scroll := c.DefaultQuery("scroll", "0")
	if paginacion, err := strconv.Atoi(scroll); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx, client := db.ElasticInit()

		categoriaQuery := elastic.NewTermQuery("title", query)
		if searchResult, errSearch := client.Search().
			Index(db.GetIndex()).Query(categoriaQuery).From(10 * paginacion).Size(10).Do(ctx); err != nil {
			c.String(http.StatusInternalServerError, errSearch.Error())
		} else {
			tweets := make([]models.ProductIndex, len(searchResult.Hits.Hits))
			for i, item := range searchResult.Each(reflect.TypeOf(models.ProductIndex{})) {
				tweetRes := item.(models.ProductIndex)
				tweets[i] = tweetRes
			}

			if len(searchResult.Hits.Hits) > 0 {
				c.JSON(http.StatusOK, tweets)
			} else {
				c.String(http.StatusNotFound, "Not Found")
			}
		}
	}
}
