package controllers

import (
	"login-signup-api/models"

	"github.com/gin-gonic/gin"
)


func ChangeRole(c *gin.Context) {
	var access models.ManageAccess
	c.BindJSON(&access)
	_, err := access.SaveAccess()
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, access)
}


func GetPermissions(c *gin.Context) {
	user_id := c.Param("id")
	permissions, _ := models.GetPermissions(user_id)
	c.JSON(200, permissions)
}
