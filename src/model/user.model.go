package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId        uint            `json:"user_id" gorm:"primaryKey"`
	Teams         []*Team         `json:"teams" gorm:"many2many:user_team"`
	Organizations []*Organization `json:"organizations" gorm:"many2many:user_organization"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
