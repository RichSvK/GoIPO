package repository

import (
	"IPO/configs"
	"IPO/models/entity"
	"context"
)

type BrokerRepository interface {
	Save(ctx context.Context, broker entity.Broker) error
}

type BrokerRepositoryImpl struct{}

func NewBrokerRepository() BrokerRepository {
	return &BrokerRepositoryImpl{}
}

func (repository *BrokerRepositoryImpl) Save(ctx context.Context, broker entity.Broker) error {
	db := configs.GetDatabaseInstance()
	err := db.WithContext(ctx).Create(&broker).Error
	return err
}
