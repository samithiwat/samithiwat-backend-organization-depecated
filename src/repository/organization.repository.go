package repository

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"github.com/samithiwat/samithiwat-backend-organization/src/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) FindAll(meta *proto.PaginationMetadata, perms *[]*model.Organization) error {
	return r.db.Scopes(Pagination(perms, meta)).Find(&perms).Count(&meta.ItemCount).Error
}

func (r *OrganizationRepository) FindOne(id int, perm *model.Organization) error {
	return r.db.Preload(clause.Associations).First(&perm, id).Error
}

func (r *OrganizationRepository) Create(perm *model.Organization) error {
	return r.db.Create(&perm).Error
}

func (r *OrganizationRepository) Update(id int, perm *model.Organization) error {
	return r.db.Where(id).Updates(&perm).First(&perm).Error
}

func (r *OrganizationRepository) Delete(id int, perm *model.Organization) error {
	return r.db.First(&perm, id).Delete(&model.Organization{}).Error
}
