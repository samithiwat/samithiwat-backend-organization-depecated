package repository

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) FindAll(*proto.PaginationMetadata, *[]*model.Team) error {
	return nil
}

func (r *TeamRepository) FindOne(uint, *model.Team) error {
	return nil
}

func (r *TeamRepository) Create(*model.Team) error {
	return nil
}

func (r *TeamRepository) Update(uint, *model.Team) error {
	return nil
}

func (r *TeamRepository) Delete(uint, *model.Team) error {
	return nil
}
