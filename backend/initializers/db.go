package initializers

import (
	"github.com/prezeswastaken/election/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "host=pgsql user=sail password=password dbname=election port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Couldn't connect to database!")
	} else {
		println("Connected to database succesfully!")
	}
	DB = db
}
func Migrate() {
	DB.AutoMigrate(&models.Contestant{})
}
