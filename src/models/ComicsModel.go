package models

import (
	"gorm.io/gorm"
)

type Comics struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Volume    int    `json:"volume"`
	Publisher string `json:"publisher"`
	Editor    string `json:"editor"`
}
