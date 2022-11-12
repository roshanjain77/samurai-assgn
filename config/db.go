package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
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

// Dashboard model
type Dashboard struct {
	gorm.Model
	UserId        int      `json:"user_id" gorm: "foreignKey:UserId"` // foreign key
	DashboardName string   `json:"dashboard_name" gorm:"size:255;not null"`
	Widgets       string `json:"widgets"`
}

var DB *gorm.DB

func Connect(){
	err := godotenv.Load("env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	  
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",DbHost, DbUser, DbPassword, DbName, DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil{
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	} else{
		fmt.Println("We are connected to the database ")
	}
	db.AutoMigrate(User{})
	db.AutoMigrate(Dashboard{})
	DB = db

}