package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "uyumak"
	dbname   = "zebgarden"
)

var db *gorm.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = conn
	db.Debug().AutoMigrate(Dealer{}, Admin{}, Image{}, Product{}, Order{}, Cart{}, CartOnOrder{}, ProductImage{})
	fmt.Println("Veri Tabanı bağlantısı başarılı!")
}
func GetDB() *gorm.DB {
	return db
}
