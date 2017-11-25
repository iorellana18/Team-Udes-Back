package models

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	ProductCRUD(app)
	UserCRUD(app)
	UserTypeCRUD(app)
}
