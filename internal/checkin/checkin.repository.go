package checkin

import (
	"github.com/isd-sgcu/rpkm67-checkin/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(checkin *model.Checkin) error
	FindByEmail(email string, checkins *[]*model.Checkin) error
}

type repositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{Db: db}
}

func (r *repositoryImpl) Create(checkin *model.Checkin) error {
	return r.Db.Create(checkin).Error
}

func (r *repositoryImpl) FindByEmail(email string, checkins *[]*model.Checkin) error {
	return r.Db.Where("email = ?", email).Find(&checkins).Error
}
