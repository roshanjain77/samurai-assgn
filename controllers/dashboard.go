package controllers

import (
	"login-signup-api/config"
	"login-signup-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardInput struct {
	DashboardName string `json:"dashboard_name"`
	Widgets       string `json:"widgets"`
}

func CreateDashboard(c *gin.Context) {
	var input DashboardInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("id")
	userId, _ := strconv.Atoi(userID)

	var dashboard models.Dashboard
	dashboard.DashboardName = input.DashboardName
	dashboard.Widgets = input.Widgets
	dashboard.UserId = userId

	// if no such user exists
	if err := config.DB.Where("id = ?", userID).First(&models.User{}).Error; err != nil {
		c.JSON(400, gin.H{"error": "No such user exists"})
		return
	} 

	_, err2 := dashboard.SaveDashboard()

	if err2 != nil {
		c.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(200, gin.H{"data": input})
}

// get all dashboards of a user
func GetDashboards(c *gin.Context) {
	var dashboards []models.Dashboard
	userID := c.Param("id")
	config.DB.Where("user_id = ?", userID).Find(&dashboards)
	c.JSON(200, gin.H{"data": dashboards})
}

// add a widget to a dashboard
func AddWidget(c *gin.Context) {
	// get dashboard id
	dashboardID := c.Param("id")

	// get widget id
	widget := c.Param("widget")

	// check if widget exists in models.widgets_mapping
	if _, ok := models.WidgetsMapping[widget]; !ok {
		c.JSON(400, gin.H{"error": "No such widget exists"})
		return
	}

	// get dashboard
	var dashboard models.Dashboard
	config.DB.Where("id = ?", dashboardID).First(&dashboard)

	// check if dashboard exists
	if dashboard.ID == 0 {
		c.JSON(400, gin.H{"error": "No such dashboard exists"})
		return
	}

	// add widget to dashboard
	dashboard.Widgets = dashboard.Widgets + ":" + widget

	// update dashboard
	config.DB.Save(&dashboard)

	c.JSON(200, gin.H{"data": dashboard})
}
