package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
