package db

import (
	elastic "gopkg.in/olivere/elastic.v5"

	"context"
	"log"

	"github.com/citiaps/aysana-backend/utils"
)

var (
	host     string
	port     string
	username string
	pass     string
	name     string
	index    string
	typeData string
)

func ElasticSetup() {
	host = utils.Config.Elastic.Host
	port = utils.Config.Elastic.Port
	username = utils.Config.Elastic.Username
	pass = utils.Config.Elastic.Pass
	name = utils.Config.Elastic.Name
	index = utils.Config.Elastic.Index
	typeData = utils.Config.Elastic.Type
}

func createIndex(ctx context.Context, client *elastic.Client) {
	exists, err := client.IndexExists(index).Do(ctx)
	utils.Check(err)

	if !exists {
		body := "{\"mappings\": {\""+index+"\": {\"properties\": {\"Texto\": {\"type\":     \"text\", \"fielddata\": true }}}}}"
		createIndex, err := client.CreateIndex(index).Body(body).Do(ctx)
		utils.Check(err)
		if !createIndex.Acknowledged {
			log.Fatal("Not Acknowledged")
		}
	}
}

func ElasticInit() (context.Context, *elastic.Client) {
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL("http://"+host+":"+port), elastic.SetBasicAuth(username, pass))
	utils.Check(err)

	createIndex(ctx, client)

	return ctx, client
}

func ElasticIndex(ctx context.Context, client *elastic.Client, data interface{}) *elastic.IndexResponse {
	put, err := client.Index().
		Index(index).
		Type(typeData).
		BodyJson(data).
		Do(ctx)
	utils.Check(err)

	return put
}

func GetIndex() string{
	return index
}
