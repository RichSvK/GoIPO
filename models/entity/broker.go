package entity

type Broker struct {
	Broker_Code string `gorm:"type:char(2);primary_key"`
	Broker_Name string `gorm:"type:VARCHAR(41)"`

	// Relationship
	Detail_IPO []IPO_Detail `gorm:"foreignKey:uw_code;references:broker_code"`
}

// Make table name from default "brokers" to "broker_underwriter"
func (broker *Broker) TableName() string {
	return "broker_underwriter"
}
