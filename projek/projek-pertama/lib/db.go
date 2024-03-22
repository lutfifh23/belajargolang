package lib

import (

	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	connectionString := "host=localhost port=5432 user=postgres password=#Cancer23 dbname=sesi7 sslmode=disable"
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
