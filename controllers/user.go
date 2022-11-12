package controllers

import (
	"fmt"
	"login-signup-api/config"
	"login-signup-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var login LoginInput
	var user models.User
	c.BindJSON(&login)
	config.DB.Where("email = ?", login.Email).Find(&user)
	fmt.Println("logged in user: ", user)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "user does not exist!"})
		return
	}
	if login.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Incorrect Password!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func FindUserByEmail(c *gin.Context) {
	var user models.User
	fmt.Println(c.Param("email"))
	config.DB.Where("email = ?", c.Param("email")).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "user does not exist!"})
		return
	}
	c.JSON(200, &user)
}

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	fmt.Println(user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Find(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
