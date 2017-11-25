package models

import (
	"github.com/gin-gonic/gin"
	"github.com/iorellana18/Team-Udes-Back/db"
	"net/http"
)

type User_Types struct {
	Id   int    `gorm:"column:id;not null;" json:"id"`
	Type string `gorm:"column:type;not null;" json:"type"`
}

func UserTypeCRUD(app *gin.Engine) {

	app.GET("/userType/:id", UserTypeFetchOne)
	app.GET("/userType/", UserTypeFetchAll)
	app.POST("/userType/", UserTypeCreate)
	app.PUT("/userType/:id", UserTypeUpdate)
	app.DELETE("/userType/:id", UserTypeRemove)

}

func UserTypeFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var userType User_Types
	if err := db.Find(&userType, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, userType)
	}
}

func UserTypeFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var userTypes []User_Types
	if err := db.Find(&userTypes).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, userTypes)
	}
}

func UserTypeCreate(c *gin.Context) {
	var userType User_Types
	if err := c.BindJSON(&userType); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		db := db.Database()
		defer db.Close()

		if err := db.Create(&userType).Error; err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusCreated, userType)
		}
	}

}

func UserTypeUpdate(c *gin.Context) {
	var userType User_Types
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&userType).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		if err := c.BindJSON(&userType); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		} else {
			db.Save(&userType)
			c.JSON(http.StatusOK, userType)
		}
	}
}

func UserTypeRemove(c *gin.Context) {
	var userType User_Types

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&userType).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&userType)
		c.JSON(http.StatusOK, userType)
	}
}
