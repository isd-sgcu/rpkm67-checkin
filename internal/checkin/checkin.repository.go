package checkin

import (
	"context"

	"github.com/isd-sgcu/rpkm67-model/model"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, checkIn *model.CheckIn) error
	FindByEmail(ctx context.Context, email string, checkIns *[]*model.CheckIn) error
	FindByUserId(ctx context.Context, userId string, checkIns *[]*model.CheckIn) error
}

type repositoryImpl struct {
	Db     *gorm.DB
	tracer trace.Tracer
}

func NewRepository(db *gorm.DB, tracer trace.Tracer) Repository {
	return &repositoryImpl{Db: db, tracer: tracer}
}

func (r *repositoryImpl) Create(ctx context.Context, checkIn *model.CheckIn) error {
	_, span := r.tracer.Start(ctx, "repository.checkin.Create")
	defer span.End()

	return r.Db.Create(checkIn).Error
}

func (r *repositoryImpl) FindByEmail(ctx context.Context, email string, checkIns *[]*model.CheckIn) error {
	_, span := r.tracer.Start(ctx, "repository.checkin.FindByEmail")
	defer span.End()

	return r.Db.Where("email = ?", email).Find(&checkIns).Error
}

func (r *repositoryImpl) FindByUserId(ctx context.Context, userId string, checkIns *[]*model.CheckIn) error {
	_, span := r.tracer.Start(ctx, "repository.checkin.FindByUserId")
	defer span.End()

	return r.Db.Where("user_id = ?", userId).Find(&checkIns).Error
}
