package config

import (
	"github.com/19abhishek/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// db, err := gorm.Open((postgres.Open("host=mahmud.db.elephantsql.com user=rqypkvuy password=kl4r_kszcdNYERMqEM8-Jw-38v_Rs0u- dbname=rqypkvuy port=5432 sslmode=disable")))
	db, err := gorm.Open((postgres.Open("postgres://postgres:postgres@localhost:5432/postgres")))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}