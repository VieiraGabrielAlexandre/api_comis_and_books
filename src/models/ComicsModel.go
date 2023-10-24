package models

import (
	"gorm.io/gorm"
)

type Comics struct {
	gorm.Model
	Title     string `json:"title"`
	Volume    int    `json:"volume"`
	Publisher string `json:"publisher"`
	Editor    string `json:"editor"`
}
