package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Members        []*User `json:"members" gorm:"many2many:user_team""`
	ParentID       uint    `json:"parent_id"`
	SubTeams       []*Team `json:"sub_team" gorm:"foreignkey:ParentID"`
	OrganizationID uint    `json:"organization_id"`
}
