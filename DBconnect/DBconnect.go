package dbconnection

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(dns string) (a *gorm.DB) {
	var b string = dns
	db, err := gorm.Open(mysql.Open(b), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
