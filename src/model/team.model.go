package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ParentID    uint    `json:"parent_id"`
	SubTeams    []*Team `json:"sub_team" gorm:"foreignkey:ParentID"`
}
