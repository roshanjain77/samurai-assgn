package models

import (
	"errors"
	"github.com/harranali/authority"
	"login-signup-api/config"
	"gorm.io/gorm"
)

type ManageAccess struct {
	gorm.Model
	AdminID int    `json:"admin_id" gorm: "foreignKey:UserId"` // foreign key
	UserID  int    `json:"user_id" gorm: "foreignKey:UserId"`  // foreign key
	Role    string `json:"role" gorm:"size:255;not null"`
}

// save access
func (access *ManageAccess) SaveAccess() (*ManageAccess, error) {
	
	auth := authority.Resolve()

	var admin User
	config.DB.Where("id = ?", access.AdminID).Find(&admin)

	// check if admin exists
	if admin.ID == 0 {
		return &ManageAccess{}, errors.New("Admin Id is not valid")
	}

	// check if user is admin
	if admin.Role != "admin" {
		return &ManageAccess{}, errors.New("Only Admin can grant/revoke access")
	}

	// check if user exists
	var user User
	config.DB.Where("id = ?", access.UserID).Find(&user)
	if user.ID == 0 {
		return &ManageAccess{}, errors.New("User does not exist")
	}

	// user should not be admin
	if user.Role == "admin" {
		return &ManageAccess{}, errors.New("User is admin, you cannot change admin role")
	}

	// the role should exist in auth.GetRoles()
	roles, _ := auth.GetRoles()
	for _, role := range roles {
		if role == access.Role {
			err := config.DB.Create(&access).Error
			user.Role = access.Role
			config.DB.Save(&user)
			if err != nil {
				return &ManageAccess{}, err
			}
			return access, nil
		}
	}

	return &ManageAccess{}, errors.New("Role does not exist")
}


func GetPermissions(userId string) ([]string, error) {
	var user User
	config.DB.Where("id = ?", userId).Find(&user)
	if user.ID == 0 {
		return nil, errors.New("User does not exist")
	}

	auth := authority.Resolve()
	permissions, err := auth.GetPermissions()

	user_permissions := []string{}
	for _, permission := range permissions {
		ok, _ := auth.CheckRolePermission(user.Role, permission)
		if ok {
			user_permissions = append(user_permissions, permission)
		}

	}

	return user_permissions, err

}
