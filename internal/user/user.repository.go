package user

import (
	"github.com/isd-sgcu/rpkm67-model/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(email string, user *model.User) error
}

type repositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{Db: db}
}

func (r *repositoryImpl) FindByEmail(email string, user *model.User) error {
	return r.Db.First(user, "email = ?", email).Error
}
