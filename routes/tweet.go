package routes

import (
	elastic "gopkg.in/olivere/elastic.v5"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/models"
	"github.com/iorellana18/Team-Udes-Back/utils"
)

func TweetsPorCategoria(c *gin.Context) {
	categoria := c.Query("id")
	idCategoria, err := strconv.Atoi(categoria)
	utils.Check(err)

	scroll := c.DefaultQuery("scroll", "0")
	paginacion, err := strconv.Atoi(scroll)
	utils.Check(err)

	ctx, client := db.ElasticInit()

	categoriaQuery := elastic.NewTermQuery("Categorias", idCategoria)
	searchResult, errSearch := client.Search().
		Index(db.GetIndex()).Query(categoriaQuery).
		Sort("CreateAt", true).From(10 * paginacion).Size(10).Do(ctx)
	utils.Check(errSearch)

	tweets := make([]models.TweetShort, len(searchResult.Hits.Hits))
	for i, item := range searchResult.Each(reflect.TypeOf(models.TweetShort{})) {
		tweetRes := item.(models.TweetShort)
		tweets[i] = tweetRes
	}

	if len(searchResult.Hits.Hits) > 0 {
		c.JSON(http.StatusOK, tweets)
	} else {
		c.String(http.StatusNotFound, "Not Found")
	}
}

func CantTweetsPorCategorias(c *gin.Context) {
	ctx, client := db.ElasticInit()

	termsAggregation := elastic.NewTermsAggregation().Field("Categorias")

	query, errSearch := client.Search(db.GetIndex()).Aggregation("CountCategorias", termsAggregation).Do(ctx)
	utils.Check(errSearch)

	bucket, _ := query.Aggregations.BucketScript("CountCategorias")

	var docCount DocCount
	if bytesJson, err := bucket.Aggregations["buckets"].MarshalJSON(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		if err := json.Unmarshal(bytesJson, &docCount); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, docCount)
		}
	}
}
