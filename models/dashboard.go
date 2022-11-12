package models

import (
	"errors"
	"fmt"
	"login-signup-api/config"
	"strings"

	"gorm.io/gorm"
)

// All widgets
var WidgetsMapping = map[string]string{
	"login":    "Login",
	"signup":   "Signup",
	"forgot":   "Forgot",
	"reset":    "Reset",
	"verify":   "Verify",
	"logout":   "Logout",
	"settings": "Settings",
}

type Dashboard struct {
	gorm.Model
	UserId        int    `json:"user_id" gorm: "foreignKey:UserId"` // foreign key
	DashboardName string `json:"dashboard_name" gorm:"size:255;not null"`
	Widgets       string `json:"widgets" gorm:"default:null"`
}

// save dashboard
func (dashboard *Dashboard) SaveDashboard() (*Dashboard, error) {
	// split widgets into array colon separated
	widgets := strings.Split(dashboard.Widgets, ":")

	fmt.Println(widgets)
	// check if widgets are valid
	for _, widget := range widgets {
		if _, ok := WidgetsMapping[widget]; !ok {
			return nil, errors.New("Invalid widget")
		}
	}

	err := config.DB.Create(&dashboard).Error
	if err != nil {
		return &Dashboard{}, err
	}
	return dashboard, nil
}
