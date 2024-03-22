package model

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address *string `json:"address"`
	UUID    string  `gorm:"primaryKey"`
	// Cards     []CreditCard
	// DeletedAt gorm.DeletedAt
	// Status  bool    `json:"status"`
}

func (p *Person) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("halo ini dari hook before create person")
	p.UUID = uuid.NewString()
	// p.Cards = append(p.Cards, CreditCard{
	// 	CardNumber: "XYZ-123",
	// })
	return nil
}
