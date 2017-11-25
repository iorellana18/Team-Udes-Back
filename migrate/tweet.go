package migrate

import (
	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/models"

	"gopkg.in/mgo.v2/bson"

	"strconv"
)

func getAllTweets() []models.Tweet {
	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("tweets", database)

	var tweets []models.Tweet
	if err := collection.Find(bson.M{}).All(&tweets); err != nil {
		return nil
	}

	return tweets
}

func indexTweet(tweet models.Tweet) {
	ctx, client := db.ElasticInit()

	tweetShort := models.TweetShort{strconv.FormatInt(tweet.IdTweet, 10), tweet.Text, tweet.InformationType, tweet.CodComuna, tweet.CodRegion, tweet.CreatedAt}

	db.ElasticIndex(ctx, client, tweetShort)
}

func MongoToElastic() {
	tweets := getAllTweets()

	for _, tweet := range tweets {
		indexTweet(tweet)
	}
}
