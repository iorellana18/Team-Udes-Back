package models

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/iorellana18/Team-Udes-Back/db"
)

type Users struct {
	Email     string `gorm:"column:email;not null;" json:"email"`
	Id        int    `gorm:"column:id;not null;" json:"id"`
	Name      string `gorm:"column:name;not null;" json:"name"`
	Password  string `gorm:"column:password;not null;" json:"password"`
	Lastname  string `gorm:"column:lastname;not null;" json:"lastname"`
	User_type int    `gorm:"column:user_type;not null;" json:"user_type"`
}

func UserCRUD(app *gin.Engine) {
	app.GET("/user/:id", UserFetchOne)
	app.GET("/user/", UserFetchAll)
	app.POST("/user/", UserCreate)
	app.PUT("/user/:id", UserUpdate)
	app.DELETE("/user/:id", UserRemove)
}

func UserFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var user Users
	if err := db.Find(&user, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, user)
	}
}

func UserFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var users []Users
	if err := db.Find(&users).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, users)
	}
}

func UserCreate(c *gin.Context) {
	var user Users
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		db := db.Database()
		defer db.Close()

		if err := db.Create(&user).Error; err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusCreated, user)
		}
	}

}

func UserUpdate(c *gin.Context) {
	var user Users
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		if err := c.BindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		} else {
			db.Save(&user)
			c.JSON(http.StatusOK, user)
		}
	}
}

func UserRemove(c *gin.Context) {
	var user Users

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&user)
		c.JSON(http.StatusOK, user)
	}
}
