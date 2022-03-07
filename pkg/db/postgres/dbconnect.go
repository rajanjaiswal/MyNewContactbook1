//connecting to database
package postgres

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rajanjaiswal/MyNewContactbook/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm" // ORM IS OBJECT RELATED MAPPING
)

// var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.Contact{},
		&models.Address{},
		&models.Phone{},
	)
	if err != nil {
		panic(err)
	}
	//DB = db
	return db

}
