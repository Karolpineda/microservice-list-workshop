package models

import (
	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type Workshop struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}

func GetAllWorkshops(db *gorm.DB) ([]Workshop, error) {
	var workshops []Workshop
	if err := db.Find(&workshops).Error; err != nil {
		return nil, err
	}
	return workshops, nil
}
