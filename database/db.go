package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alexkarpovich/mowy-api-go/database/regs"
	"github.com/alexkarpovich/mowy-api-go/database/users"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func InitDB() {
	params := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"))

	db, err := gorm.Open("postgres", params)
	if err != nil {
		log.Fatal(err)
	}
	//defer Db.Close()

	db.LogMode(true)
	db.AutoMigrate(&users.User{}, &regs.Registration{})

	Db = db
}
