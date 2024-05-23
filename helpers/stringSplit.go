package helpers

import (
	"IPO/models/entity"
	"strconv"
	"strings"
)

func SplitStockString(result []byte) entity.Stock {
	var stock entity.Stock
	var temp uint64 = 0
	stockData := strings.Split(string(result), ",")
	stock.StockCode = stockData[0]

	temp, _ = strconv.ParseUint(stockData[1], 10, 32)
	stock.Price = uint32(temp)
	temp, _ = strconv.ParseUint(stockData[6], 10, 32)
	stock.Nominal = uint32(temp)
	temp, _ = strconv.ParseUint(stockData[13], 10, 8)
	stock.LockUp = int8(temp)

	stock.IPO_Shares, _ = strconv.ParseUint(stockData[2], 10, 64)
	stock.ListedShares, _ = strconv.ParseUint(stockData[3], 10, 64)
	stock.Equity, _ = strconv.ParseInt(stockData[4], 10, 64)
	stock.Warrant, _ = strconv.ParseUint(stockData[5], 10, 64)
	stock.MCB, _ = strconv.ParseUint(stockData[7], 10, 64)
	stock.SubscribedStock, _ = strconv.ParseUint(stockData[14], 10, 64)

	stock.IsAffiliated, _ = strconv.ParseBool(stockData[8])
	stock.IsAcceleration, _ = strconv.ParseBool(stockData[9])
	stock.IsNew, _ = strconv.ParseBool(stockData[10])
	stock.IsFullCommitment, _ = strconv.ParseBool(stockData[11])
	stock.IsNotInvolvedCase, _ = strconv.ParseBool(stockData[12])
	return stock
}

func SplitBrokerString(result []byte) entity.Broker {
	var broker entity.Broker
	brokerData := strings.Split(string(result), ",")
	broker.Broker_Code = brokerData[0]
	broker.Broker_Name = brokerData[1]
	return broker
}

func SplitDetailString(result []byte) entity.IPO_Detail {
	var detail entity.IPO_Detail
	detailData := strings.Split(string(result), ",")
	detail.StockCode = detailData[0]
	detail.UW_Code = detailData[1]
	temp, _ := strconv.ParseUint(detailData[2], 10, 64)
	detail.UwShares = uint64(temp)
	return detail
}
