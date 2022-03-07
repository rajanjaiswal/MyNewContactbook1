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
	dsn := fmt.Sprintf("host=localhost user=postgres password=123 dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.Contact{},
		&models.Address{},
		&models.Phone{},
	)
	//DB = db
	return db

}
