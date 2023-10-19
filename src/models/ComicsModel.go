package models

import (
	"gorm.io/gorm"
)

type Comics struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Volume    int    `json:"volume"`
	Publisher string `json:"publisher"`
	Editor    string `json:"editor"`
}
