package crud

import (
	"projek-kedua/entity"

	"gorm.io/gorm"
)

func InsertProduct(db *gorm.DB, product *entity.Products) error {
	result := db.Create(product)

	if result.Error != nil {
		panic("Failed to insert order")

	}
}
