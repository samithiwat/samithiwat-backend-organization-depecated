package repository

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) FindAll(*proto.PaginationMetadata, *[]*model.Organization) error {
	return nil
}

func (r *OrganizationRepository) FindOne(uint, *model.Organization) error {
	return nil
}

func (r *OrganizationRepository) Create(*model.Organization) error {
	return nil
}

func (r *OrganizationRepository) Update(uint, *model.Organization) error {
	return nil
}

func (r *OrganizationRepository) Delete(uint, *model.Organization) error {
	return nil
}
