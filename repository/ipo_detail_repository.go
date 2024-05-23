package repository

import (
	"IPO/configs"
	"IPO/models/entity"
	"context"
)

type IpoDetailRepository interface {
	Save(ctx context.Context, detail entity.IPO_Detail) error
}

type IpoDetailRepositoryImpl struct{}

func NewDetailRepository() IpoDetailRepository {
	return &IpoDetailRepositoryImpl{}
}

func (repository *IpoDetailRepositoryImpl) Save(ctx context.Context, detail entity.IPO_Detail) error {
	db := configs.GetDatabaseInstance()
	err := db.WithContext(ctx).Create(&detail).Error
	return err
}
