package entity

type IPO_Detail struct {
	StockCode string `gorm:"type:char(4);primaryKey"`
	UW_Code   string `gorm:"type:char(2);primaryKey"`
	UwShares  uint64 `gorm:"type:BIGINT UNSIGNED"`
}

// Make table name from default "ipo_details" to "ipo_detail"
func (broker *IPO_Detail) TableName() string {
	return "ipo_detail"
}
