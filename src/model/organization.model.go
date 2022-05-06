package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Teams       []*Team `json:"teams"`
	Members     []*User `json:"members" gorm:"many2many:user_organization"`
}
