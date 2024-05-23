package helpers

import (
	"IPO/models/response"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func ExportResponse(filePath string, header string, listStock []response.StockResponse) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Fail to open file because", err.Error())
		return
	}
	defer file.Close()

	file.WriteString(header)
	for _, stock := range listStock {
		file.WriteString(stock.StockCode + ",")
		file.WriteString(strconv.FormatUint(uint64(stock.Price), 10) + ",")
		file.WriteString(strconv.FormatUint(stock.IPO_Shares, 10) + ",")
		file.WriteString(strconv.FormatUint(stock.ListedShares, 10) + ",")
		file.WriteString(strconv.FormatInt(stock.Equity, 10) + ",")

		if stock.Warrant != 0 {
			warrant := big.NewRat(int64(stock.IPO_Shares), int64(stock.Warrant))
			file.WriteString(fmt.Sprintf("\"%s : %s\",", warrant.Num().String(), warrant.Denom().String()))
		} else {
			file.WriteString("0,")
		}

		file.WriteString(strconv.FormatUint(uint64(stock.Nominal), 10) + ",")
		file.WriteString(strconv.FormatUint(stock.MCB, 10) + ",")
		file.WriteString(strconv.FormatBool(stock.IsAffiliated) + ",")
		file.WriteString(strconv.FormatBool(stock.IsAcceleration) + ",")
		file.WriteString(strconv.FormatBool(stock.IsNew) + ",")
		file.WriteString(strconv.FormatInt(int64(stock.LockUp), 10) + ",")
		file.WriteString(strconv.FormatUint(stock.SubscribedStock, 10) + ",\"{")
		WriterUnderwriter(file, stock.AllUnderwriter, stock.AllShares, float64(stock.IPO_Shares))
		file.WriteString(strconv.FormatUint(stock.Amount, 10) + "\n")
	}
}

func WriterUnderwriter(file *os.File, AllUnderwriter string, AllShares string, IPO_Shares float64) {
	underwriter := strings.Split(AllUnderwriter, ",")
	uwShares := strings.Split(AllShares, ",")
	size := len(underwriter) - 1
	var percentage float64 = 0

	for i := 0; i < size; i++ {
		share, _ := strconv.ParseFloat(uwShares[i], 64)
		percentage = share / IPO_Shares * 100
		file.WriteString(fmt.Sprintf("%s : %.2f%%, ", underwriter[i], percentage))
	}
	share, _ := strconv.ParseFloat(uwShares[size], 64)
	percentage = share / IPO_Shares * 100
	file.WriteString(fmt.Sprintf("%s : %.2f%%}\",", underwriter[size], percentage))
}
