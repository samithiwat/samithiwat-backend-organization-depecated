package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	OrganizationID uint `json:"organization_id"`
}
