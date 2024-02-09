package repo

import (
	"github.com/anjush-bhargavan/golib_admin/pkg/dom"
	inter "github.com/anjush-bhargavan/golib_admin/pkg/repo/interfaces"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) inter.AdminRepoInter {
	return &AdminRepo{
		DB: db,
	}
}

func (a *AdminRepo) FindAdmin(username string) (*dom.AdminModel, error) {
	var admin dom.AdminModel
	err := a.DB.Where("username = ?", username).First(&admin).Error
	return &admin, err
}
