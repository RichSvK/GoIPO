package repository

import (
	"IPO/configs"
	"IPO/helpers"
	"IPO/models/entity"
	"IPO/models/response"
	"context"
)

type StockRepository interface {
	Save(ctx context.Context, stock entity.Stock) error
	FindByUnderwriter(ctx context.Context, underwriter string) ([]response.StockResponse, error)
	FindByValue(ctx context.Context, value int, underwriter string) ([]response.StockResponse, error)
}

type StockRepositoryImpl struct{}

func NewStockRepository() StockRepository {
	return &StockRepositoryImpl{}
}

func (repository *StockRepositoryImpl) Save(ctx context.Context, stock entity.Stock) error {
	db := configs.GetDatabaseInstance()
	err := db.WithContext(ctx).Create(&stock).Error
	return err
}

func (repository *StockRepositoryImpl) FindByUnderwriter(ctx context.Context, underwriter string) ([]response.StockResponse, error) {
	db := configs.PoolDB
	var listStock []response.StockResponse = nil
	err := db.WithContext(ctx).
		Table("ipo_detail id").
		Select("sub.code AS stock_code, GROUP_CONCAT(id.uw_code) AS all_underwriter, GROUP_CONCAT(uw_shares) AS all_shares, price, ipo_shares, listed_shares, equity, warrant, nominal, mcb, is_affiliated, is_acceleration, is_new, lock_up, subscribed_stock, amount").
		Joins("JOIN (SELECT s.stock_code AS code, price, ipo_shares, listed_shares, equity, warrant, nominal, mcb, is_affiliated, is_acceleration, is_new, lock_up, subscribed_stock, (price * ipo_shares) AS amount, CAST(uw_shares / ipo_shares AS FLOAT) AS percentage FROM stock_ipo s JOIN ipo_detail ids ON s.stock_code = ids.stock_code WHERE ids.uw_code = ?) AS sub ON sub.code = id.stock_code", underwriter).
		Group("sub.code, percentage").
		Order("percentage DESC").
		Scan(&listStock).
		Error

	return listStock, err
}

func (repository *StockRepositoryImpl) FindByValue(ctx context.Context, value int, underwriter string) ([]response.StockResponse, error) {
	listStock := []response.StockResponse{}
	db := configs.PoolDB
	if underwriter != "ALL" {
		db = db.Joins("JOIN (SELECT s.stock_code AS stock_code, price, ipo_shares, listed_shares, equity, warrant, nominal, mcb, is_affiliated, is_acceleration, is_new, lock_up, subscribed_stock FROM ipo_detail ids JOIN stock s ON s.stock_code = ids.stock_code WHERE uw_code = ?) ts ON ts.stock_code = id.stock_code", underwriter)
	} else {
		db = db.Joins("JOIN stock_ipo s ON id.stock_code = s.stock_code")
	}

	query := "id.stock_code, price, ipo_shares, listed_shares, equity, warrant, nominal, mcb, is_affiliated, is_acceleration, is_new, lock_up, subscribed_stock, GROUP_CONCAT(uw_code) AS all_underwriter, GROUP_CONCAT(uw_shares) AS all_shares, (price * ipo_shares) AS amount"
	err := db.Table("ipo_detail id").
		WithContext(ctx).
		Select(query).
		Where(helpers.GetAmountCondition(value)).
		Group("id.stock_code").
		Find(&listStock).
		Error
	return listStock, err
}
