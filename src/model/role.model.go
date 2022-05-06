package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	RoleID         uint `json:"role_id" gorm:"primaryKey"`
	OrganizationID uint `json:"organization_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
