package response

type StockResponse struct {
	StockCode       string
	Price           uint32
	IPO_Shares      uint64
	ListedShares    uint64
	Equity          int64
	Warrant         uint64
	Nominal         uint32
	MCB             uint64
	IsAffiliated    bool
	IsAcceleration  bool
	IsNew           bool
	LockUp          int8
	SubscribedStock uint64
	AllUnderwriter  string
	AllShares       string
	Amount          uint64
}
