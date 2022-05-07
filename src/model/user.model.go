package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        uint            `json:"user_id" gorm:"index:,unique"`
	Teams         []*Team         `json:"teams" gorm:"many2many:user_team"`
	Organizations []*Organization `json:"organizations" gorm:"many2many:user_organization"`
}
