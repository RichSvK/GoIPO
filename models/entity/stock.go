package entity

type Stock struct {
	StockCode         string `gorm:"type:char(4);primaryKey"`
	Price             uint32 `gorm:"type:BIGINT UNSIGNED"`
	IPO_Shares        uint64 `gorm:"type:BIGINT UNSIGNED"`
	ListedShares      uint64 `gorm:"type:BIGINT UNSIGNED"`
	Equity            int64  `gorm:"column:equity"`
	Warrant           uint64 `gorm:"type:BIGINT UNSIGNED"`
	Nominal           uint32 `gorm:"type: INT UNSIGNED"`
	MCB               uint64 `gorm:"type:BIGINT UNSIGNED"`
	IsAffiliated      bool
	IsAcceleration    bool
	IsNew             bool
	IsFullCommitment  bool
	IsNotInvolvedCase bool
	LockUp            int8
	SubscribedStock   uint64 `gorm:"type:BIGINT UNSIGNED"`

	// Relationship
	Detail []IPO_Detail `gorm:"foreignKey:stock_code;references:stock_code"`
}

// Make table name from default "stocks" to "stock_ipo"
func (detail *Stock) TableName() string {
	return "stock_ipo"
}
