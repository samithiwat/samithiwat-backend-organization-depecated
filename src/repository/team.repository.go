package repository

import (
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) FindAll(pagination *model.TeamPagination) error {
	return r.db.Scopes(Pagination(&pagination.Items, &pagination.Meta, r.db)).Find(&pagination.Items).Count(&pagination.Meta.ItemCount).Error
}

func (r *TeamRepository) FindOne(id uint, perm *model.Team) error {
	return r.db.Preload(clause.Associations).First(&perm, id).Error
}

func (r *TeamRepository) FindMulti(ids []uint32, teams *[]*model.Team) error {
	return r.db.Where("id IN ?", ids).Find(&teams).Error
}

func (r *TeamRepository) Create(perm *model.Team) error {
	return r.db.Create(&perm).Error
}

func (r *TeamRepository) Update(id uint, perm *model.Team) error {
	return r.db.Where(id).Updates(&perm).First(&perm).Error
}

func (r *TeamRepository) Delete(id uint, perm *model.Team) error {
	return r.db.First(&perm, id).Delete(&model.Team{}).Error
}
