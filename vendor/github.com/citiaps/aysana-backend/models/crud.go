package models

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	CategoriaCRUD(app)
	TweetCRUD(app)
}
