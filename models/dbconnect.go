//connecting to database
package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm" // ORM IS OBJECT RELATED MAPPING
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=123 dbname=MyNewContact port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&Contact{},
		&Address{},
		&Phone{},
	)
	DB = db

}
