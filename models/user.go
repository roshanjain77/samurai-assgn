package models

import (
	"fmt"
	"login-signup-api/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;" json:"name"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Role	string	`gorm:"size:255;not null;default:user" json:"role"`
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// SaveUser is a function to save a user
func SaveUser(user *User) (*User, error) {

	fmt.Println(user)
	err := config.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) FindUserByEmail(email string) (*User, error) {
	var err error
	err = config.DB.Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.ErrRecordNotFound == err {
		return &User{}, err
	}
	return user, err
}


func (user *User) UpdateUserRole(uid uint32, role string) (*User, error) {

	config.DB = config.DB.Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"role":  role,
		},
	)
	if config.DB.Error != nil {
		return &User{}, config.DB.Error
	}
	// This is the display the updated user
	err := config.DB.Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// func (user *User) FindAllUsers() (*[]User, error) {
// 	var err error
// 	users := []User{}
// 	err = DB.Model(&User{}).Limit(100).Find(&users).Error
// 	if err != nil {
// 		return &[]User{}, err
// 	}
// 	return &users, nil
// }

// func (user *User) FindUserByID(uid uint32) (*User, error) {
// 	var err error
// 	err = DB.Model(User{}).Where("id = ?", uid).Take(&user).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	if gorm.ErrRecordNotFound == err {
// 		return &User{}, err
// 	}
// 	return user, err
// }


// func (user *User) DeleteAUser(uid uint32) (int64, error) {

// 	db := DB.Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

// 	if db.Error != nil {
// 		if gorm.IsRecordNotFoundError(db.Error) {
// 			return 0, errors.New("User not found")
// 		}
// 		return 0, db.Error
// 	}
// 	return db.RowsAffected, nil
// }

