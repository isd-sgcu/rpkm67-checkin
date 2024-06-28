package checkin

import (
	"github.com/isd-sgcu/rpkm67-model/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(checkIn *model.CheckIn) error
	FindByEmail(email string, checkIns *[]*model.CheckIn) error
	FindByUserId(userId string, checkIns *[]*model.CheckIn) error
}

type repositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{Db: db}
}

func (r *repositoryImpl) Create(checkIn *model.CheckIn) error {
	return r.Db.Create(checkIn).Error
}

func (r *repositoryImpl) FindByEmail(email string, checkIns *[]*model.CheckIn) error {
	return r.Db.Where("email = ?", email).Find(&checkIns).Error
}

func (r *repositoryImpl) FindByUserId(userId string, checkIns *[]*model.CheckIn) error {
	return r.Db.Where("user_id = ?", userId).Find(&checkIns).Error
}
