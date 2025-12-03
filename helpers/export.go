package helpers

import (
	"IPO/models/response"
	"bufio"
	"fmt"
	"io"
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

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error open file:", err)
		}
	}()

	writer := bufio.NewWriter(file)

	write := func(s string) {
		if _, err := writer.WriteString(s); err != nil {
			fmt.Println("Write error:", err)
		}
	}

	write(header)

	for _, stock := range listStock {
		write(stock.StockCode + ",")
		write(strconv.FormatUint(uint64(stock.Price), 10) + ",")
		write(strconv.FormatUint(stock.IPO_Shares, 10) + ",")
		write(strconv.FormatUint(stock.ListedShares, 10) + ",")
		write(strconv.FormatInt(stock.Equity, 10) + ",")

		if stock.Warrant != 0 {
			warrant := big.NewRat(int64(stock.IPO_Shares), int64(stock.Warrant))
			write(fmt.Sprintf("\"%s : %s\",", warrant.Num().String(), warrant.Denom().String()))
		} else {
			write("0,")
		}

		write(strconv.FormatUint(uint64(stock.Nominal), 10) + ",")
		write(strconv.FormatUint(stock.MCB, 10) + ",")
		write(strconv.FormatBool(stock.IsAffiliated) + ",")
		write(strconv.FormatBool(stock.IsAcceleration) + ",")
		write(strconv.FormatBool(stock.IsNew) + ",")
		write(strconv.FormatInt(int64(stock.LockUp), 10) + ",")
		write(strconv.FormatUint(stock.SubscribedStock, 10) + ",\"{")
		WriterUnderwriter(writer, stock.AllUnderwriter, stock.AllShares, float64(stock.IPO_Shares))
		write(strconv.FormatUint(stock.Amount, 10) + "\n")
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Flush error:", err)
		return
	}
}

func WriterUnderwriter(writer io.Writer, AllUnderwriter string, AllShares string, IPO_Shares float64) {
	underwriter := strings.Split(AllUnderwriter, ",")
	uwShares := strings.Split(AllShares, ",")
	size := len(underwriter) - 1

	write := func(s string) {
		if _, err := writer.Write([]byte(s)); err != nil {
			fmt.Println("Write error:", err)
		}
	}

	for i := 0; i < size; i++ {
		share, _ := strconv.ParseFloat(uwShares[i], 64)
		percentage := share / IPO_Shares * 100
		write(fmt.Sprintf("%s : %.2f%%, ", underwriter[i], percentage))
	}
	share, _ := strconv.ParseFloat(uwShares[size], 64)
	percentage := share / IPO_Shares * 100
	write(fmt.Sprintf("%s : %.2f%%}\",", underwriter[size], percentage))
}
