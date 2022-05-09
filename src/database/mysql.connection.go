package database

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-organization/src/config"
	"github.com/samithiwat/samithiwat-backend-organization/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(conf *config.Database) (gormDb *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", conf.User, conf.Password, conf.Host, conf.Name)

	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = gormDb.AutoMigrate(model.Team{}, model.Organization{}, model.Role{}, model.Location{}, model.Contact{})
	if err != nil {
		return
	}

	return
}
