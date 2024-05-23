package helpers

func GetAmountCondition(value int) string {
	switch value {
	case 1:
		return "price * ipo_shares <= 250000000000"
	case 2:
		return "price * ipo_shares > 250000000000 AND price * ipo_shares <= 500000000000"
	case 3:
		return "price * ipo_shares > 500000000000 AND price * ipo_shares <= 1000000000000"
	case 4:
		return "price * ipo_shares > 1000000000000"
	}
	return ""
}
