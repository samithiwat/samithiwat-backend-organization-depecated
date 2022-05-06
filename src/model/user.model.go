package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Teams         []*Team         `json:"teams" gorm:"many2many:user_team"`
	Organizations []*Organization `json:"organizations" gorm:"many2many:user_organization"`
}
