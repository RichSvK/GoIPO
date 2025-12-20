package helpers

import (
	"IPO/models/entity"
	"strconv"
	"strings"
)

func SplitStockString(result []string) entity.Stock {
	var stock entity.Stock
	var temp uint64

	stock.StockCode = result[0]

	temp, _ = strconv.ParseUint(result[1], 10, 32)
	stock.Price = uint32(temp)
	temp, _ = strconv.ParseUint(result[6], 10, 32)
	stock.Nominal = uint32(temp)
	temp, _ = strconv.ParseUint(result[13], 10, 8)
	stock.LockUp = int8(temp)

	stock.IPO_Shares, _ = strconv.ParseUint(result[2], 10, 64)
	stock.ListedShares, _ = strconv.ParseUint(result[3], 10, 64)
	stock.Equity, _ = strconv.ParseInt(result[4], 10, 64)
	stock.Warrant, _ = strconv.ParseUint(result[5], 10, 64)
	stock.MCB, _ = strconv.ParseUint(result[7], 10, 64)
	stock.SubscribedStock, _ = strconv.ParseUint(result[14], 10, 64)

	stock.IsAffiliated, _ = strconv.ParseBool(result[8])
	stock.IsAcceleration, _ = strconv.ParseBool(result[9])
	stock.IsNew, _ = strconv.ParseBool(result[10])
	stock.IsFullCommitment, _ = strconv.ParseBool(result[11])
	stock.IsNotInvolvedCase, _ = strconv.ParseBool(result[12])
	return stock
}

func SplitBrokerString(result []byte) entity.Broker {
	var broker entity.Broker
	brokerData := strings.Split(string(result), ",")
	broker.Broker_Code = brokerData[0]
	broker.Broker_Name = brokerData[1]
	return broker
}

func SplitDetailString(result []string) entity.IPO_Detail {
	var detail entity.IPO_Detail

	detail.StockCode = result[0]
	detail.UW_Code = result[1]
	temp, _ := strconv.ParseUint(result[2], 10, 64)
	detail.UwShares = uint64(temp)
	return detail
}
