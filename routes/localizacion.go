package routes

import (
	elastic "gopkg.in/olivere/elastic.v5"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"net/http"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/utils"
)

type DocCount []Count

type Count struct {
	Key       int `json:"key"`
	Doc_count int `json:"doc_count"`
}

func RegionPorCategoria(c *gin.Context) {
	categoria := c.Query("id")

	ctx, client := db.ElasticInit()

	categoriaQuery := elastic.NewTermQuery("Categorias", categoria)
	termsAggregation := elastic.NewTermsAggregation().Field("CodRegion")

	query, errSearch := client.Search(db.GetIndex()).Query(categoriaQuery).Aggregation("Region", termsAggregation).Do(ctx)
	utils.Check(errSearch)

	bucket, _ := query.Aggregations.BucketScript("Region")

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

func ComunaPorCategoria(c *gin.Context) {
	categoria := c.Query("id")

	ctx, client := db.ElasticInit()

	categoriaQuery := elastic.NewTermQuery("Categorias", categoria)
	termsAggregation := elastic.NewTermsAggregation().Field("CodComuna")

	query, errSearch := client.Search(db.GetIndex()).Query(categoriaQuery).Aggregation("Comuna", termsAggregation).Do(ctx)
	utils.Check(errSearch)

	bucket, _ := query.Aggregations.BucketScript("Comuna")

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
