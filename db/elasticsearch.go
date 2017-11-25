package db

import (
	elastic "gopkg.in/olivere/elastic.v5"

	"context"
	"log"

	"github.com/iorellana18/Team-Udes-Back/utils"
)

var (
	hostElastic string
	portElastic string
	name        string
	index       string
	typeData    string
)

func ElasticSetup() {
	hostElastic = utils.Config.Elastic.Host
	portElastic = utils.Config.Elastic.Port
	name = utils.Config.Elastic.Name
	index = utils.Config.Elastic.Index
	typeData = utils.Config.Elastic.Type
}

func createIndex(ctx context.Context, client *elastic.Client) {
	exists, err := client.IndexExists(index).Do(ctx)
	utils.Check(err)

	if !exists {
		body := "{\"mappings\": {\"" + index + "\": {\"properties\": {\"Title\": {\"type\":     \"text\", \"fielddata\": true }}}}}"
		createIndex, err := client.CreateIndex(index).Body(body).Do(ctx)
		utils.Check(err)
		if !createIndex.Acknowledged {
			log.Fatal("Not Acknowledged")
		}
	}
}

func ElasticInit() (context.Context, *elastic.Client) {
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL("http://" + hostElastic + ":" + portElastic))
	utils.Check(err)

	createIndex(ctx, client)

	return ctx, client
}

func ElasticIndex(ctx context.Context, client *elastic.Client, data interface{}) *elastic.IndexResponse {
	put, err := client.Index().Index(index).Type(typeData).
		BodyJson(data).
		Do(ctx)
	utils.Check(err)

	return put
}

func GetIndex() string {
	return index
}
