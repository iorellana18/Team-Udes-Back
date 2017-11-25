package api

import (
	"github.com/gin-gonic/gin"

	"github.com/kaneshin/pigeon"
	"github.com/kaneshin/pigeon/credentials"

	"encoding/json"
	"github.com/iorellana18/Team-Udes-Back/models"
	"github.com/iorellana18/Team-Udes-Back/search"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Imagen []struct {
	LabelAnnotations []struct {
		Description string  `json:"description"`
		Mid         string  `json:"mid"`
		Score       float64 `json:"score"`
	} `json:"labelAnnotations"`
}

func AnalyzeImagen(c *gin.Context) {
	if file, err := c.FormFile("file"); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		src, err := file.Open()
		defer src.Close()

		fileTmp, err := ioutil.TempFile("img/", "img-tmp")
		defer os.Remove(fileTmp.Name())

		io.Copy(fileTmp, src)

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {

			creds := credentials.NewApplicationCredentials("config/credentials.json")
			config := pigeon.NewConfig().WithCredentials(creds)

			if client, err := pigeon.New(config); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			} else {
				feature := pigeon.NewFeature(pigeon.LabelDetection)
				if batch, err := client.NewBatchAnnotateImageRequest([]string{fileTmp.Name()}, feature); err != nil {
					c.String(http.StatusInternalServerError, err.Error())
				} else {
					if res, err := client.ImagesService().Annotate(batch).Do(); err != nil {
						c.String(http.StatusInternalServerError, err.Error())
					} else {
						body, _ := json.MarshalIndent(res.Responses, "", "  ")
						var imagens Imagen
						if err := json.Unmarshal(body, &imagens); err != nil {
							c.String(http.StatusInternalServerError, err.Error())
						} else {
							var products []models.ProductIndex
							for _, imagen := range imagens {
								for _, label := range imagen.LabelAnnotations {
									if text, err := translateText(label.Description); err != nil {
										c.String(http.StatusInternalServerError, err.Error())
									} else {
										log.Println("[Google Cloud Vision] Texto a buscar: " + text)
										if productsRes, err := search.QueryIndex(c, text, 0); err != nil {
											c.String(http.StatusInternalServerError, err.Error())
										} else {
											products = append(products, productsRes...)
										}
									}
								}
							}
							c.JSON(http.StatusOK, products)
						}
					}
				}
			}

		}
	}
}
