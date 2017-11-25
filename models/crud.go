package models

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	UserCRUD(app)
	UserTypeCRUD(app)
}
