package repository

import (
	"github.com/cde/model"
	"gorm.io/gorm"
)

var AdminRepository = newAdminRepository()

func newAdminRepository() *adminRepository {
	return &adminRepository{}
}

type adminRepository struct {
}

func (r *adminRepository) Create(db *gorm.DB, t *model.Admin) (err error) {
	err = db.Create(t).Error
	return err
}

func (r *adminRepository) Get(db *gorm.DB, username string) *model.Admin {
	ret := &model.Admin{}
	if err := db.First(ret, "username = ?", username).Error; err != nil {
		return nil
	}
	return ret
}

func (r *adminRepository) Take(db *gorm.DB, where ...interface{}) *model.Admin {
	ret := &model.Admin{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *adminRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Admin{}).Where("id = ?", id).Updates(columns).Error
	return
}