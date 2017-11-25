package models

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"net/http"
	"strconv"
	"time"

	"fmt"

	"github.com/citiaps/aysana-backend/db"
	"github.com/citiaps/aysana-backend/utils"
)

type TweetShort struct {
	Id         string
	Texto      string
	Categorias []int
	CodComuna  int
	CodRegion  int
	CreateAt   time.Time
}

type Tweet struct {
	Id                bson.ObjectId `json:"id" bson:"_id"`
	Contributors      interface{}   `json:"contributors"`
	Truncated         bool          `json:"truncated"`
	Text              string        `json:"text"`
	IsQuoteStatus     bool          `json:"is_quote_status"`
	InReplyToStatusID int64         `json:"in_reply_to_status_id"`
	IdTweet           int64         `json:"idTweet" bson:"id"`
	FavoriteCount     int           `json:"favorite_count"`
	Source            string        `json:"source"`
	Retweeted         bool          `json:"retweeted"`
	Coordinates       interface{}   `json:"coordinates"`
	Entities          struct {
		Symbols      []interface{} `json:"symbols"`
		UserMentions []struct {
			ID         int    `json:"id"`
			Indices    []int  `json:"indices"`
			IDStr      string `json:"id_str"`
			ScreenName string `json:"screen_name"`
			Name       string `json:"name"`
		} `json:"user_mentions"`
		Hashtags []struct {
			Indices []int  `json:"indices"`
			Text    string `json:"text"`
		} `json:"hashtags"`
		Urls []interface{} `json:"urls"`
	} `json:"entities"`
	InReplyToScreenName string `json:"in_reply_to_screen_name"`
	InReplyToUserID     int    `json:"in_reply_to_user_id"`
	RetweetCount        int    `json:"retweet_count"`
	IDStr               string `json:"id_str"`
	Favorited           bool   `json:"favorited"`
	User                struct {
		FollowRequestSent              bool   `json:"follow_request_sent"`
		HasExtendedProfile             bool   `json:"has_extended_profile"`
		ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
		DefaultProfileImage            bool   `json:"default_profile_image"`
		ID                             int    `json:"id"`
		ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
		Verified                       bool   `json:"verified"`
		ProfileTextColor               string `json:"profile_text_color"`
		ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
		ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
		Entities                       struct {
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		FollowersCount            int         `json:"followers_count"`
		ProfileSidebarBorderColor string      `json:"profile_sidebar_border_color"`
		IDStr                     string      `json:"id_str"`
		ProfileBackgroundColor    string      `json:"profile_background_color"`
		ListedCount               int         `json:"listed_count"`
		IsTranslationEnabled      bool        `json:"is_translation_enabled"`
		UtcOffset                 interface{} `json:"utc_offset"`
		StatusesCount             int         `json:"statuses_count"`
		Description               string      `json:"description"`
		FriendsCount              int         `json:"friends_count"`
		Location                  string      `json:"location"`
		ProfileLinkColor          string      `json:"profile_link_color"`
		ProfileImageURL           string      `json:"profile_image_url"`
		Following                 bool        `json:"following"`
		GeoEnabled                bool        `json:"geo_enabled"`
		ProfileBannerURL          string      `json:"profile_banner_url"`
		ProfileBackgroundImageURL string      `json:"profile_background_image_url"`
		ScreenName                string      `json:"screen_name"`
		Lang                      string      `json:"lang"`
		ProfileBackgroundTile     bool        `json:"profile_background_tile"`
		FavouritesCount           int         `json:"favourites_count"`
		Name                      string      `json:"name"`
		Notifications             bool        `json:"notifications"`
		URL                       interface{} `json:"url"`
		CreatedAt                 time.Time   `json:"created_at" bson:"created_at"`
		ContributorsEnabled       bool        `json:"contributors_enabled"`
		TimeZone                  interface{} `json:"time_zone"`
		Protected                 bool        `json:"protected"`
		DefaultProfile            bool        `json:"default_profile"`
		IsTranslator              bool        `json:"is_translator"`
	} `json:"user"`
	Geo                  interface{} `json:"geo"`
	InReplyToUserIDStr   string      `json:"in_reply_to_user_id_str"`
	Lang                 string      `json:"lang"`
	CreatedAt            time.Time   `json:"created_at" bson:"created_at"`
	InReplyToStatusIDStr string      `json:"in_reply_to_status_id_str"`
	Place                interface{} `json:"place"`
	InformationType      []int       `json:"information_type" bson:"information_type"`
	CodComuna            int         `json:"cod_comuna" bson:"cod_comuna"`
	CodRegion            int         `json:"cod_region" bson:"cod_region"`
}

func TweetCRUD(app *gin.Engine) {
	app.GET("/tweet/id/:id", TweetFetchID)
	app.GET("/tweet/skip/:id", TweetFetchOne)
	app.GET("/tweet/", TweetFetchAll)
	app.POST("/tweet/", TweetCreate)
	app.PUT("/tweet/:id", TweetUpdate)
	app.DELETE("/tweet/:id", TweetRemove)
}

func TweetFetchID(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.String(http.StatusNotFound, "Not ID Hex")
	}

	oid := bson.ObjectIdHex(id)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("tweets", database)

	var tweet Tweet
	if err := collection.FindId(oid).One(&tweet); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		fmt.Println(tweet.IdTweet)
		c.JSON(http.StatusOK, tweet)
	}
}

func TweetFetchOne(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	utils.Check(err)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("tweets", database)

	var tweet Tweet
	if err := collection.Find(bson.M{}).Skip(num).One(&tweet); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, tweet)
	}
}

func TweetFetchAll(c *gin.Context) {
	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("tweets", database)

	var tweet []Tweet
	if err := collection.Find(bson.M{}).All(&tweet); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, tweet)
	}
}

func TweetCreate(c *gin.Context) {
	var tweet Tweet
	err := c.BindJSON(&tweet)
	utils.Check(err)

	session := db.MongoSession()
	defer session.Close()

	database := db.MongoDatabase(session)
	collection := db.MongoCollection("tweets", database)

	if err := collection.Insert(&tweet); err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, tweet)
	}
}

func TweetUpdate(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}

func TweetRemove(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}
