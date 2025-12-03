package services

import (
	"IPO/helpers"
	"IPO/repository"
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

type StockService interface {
	InsertStock(fileName string)
	ExportByUnderwriter(underwriter string)
	ExportByValue(value int, underwriter string)
}

type StockServiceImpl struct {
	Stock_Repository repository.StockRepository
}

func NewStockService(stockRepository repository.StockRepository) StockService {
	return &StockServiceImpl{
		Stock_Repository: stockRepository,
	}
}

func (service *StockServiceImpl) InsertStock(fileName string) {
	if fileName == "" {
		return
	}

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	reader := bufio.NewReader(file)

	// Remove Header
	_, _, err = reader.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for {
		result, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		stock := helpers.SplitStockString(result)

		err = service.Stock_Repository.Save(ctx, stock)
		if err != nil {
			fmt.Println("Error insert stock data")
			return
		}
	}
	fmt.Println("Success Insert Data")
}

func (service *StockServiceImpl) ExportByUnderwriter(underwriter string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	listStock, err := service.Stock_Repository.FindByUnderwriter(ctx, underwriter)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(listStock) == 0 {
		fmt.Printf("No stock with %s underwriter\n", underwriter)
		return
	}

	filePath := "output/UW/UW_" + underwriter + ".csv"
	header := "Stock Code,price,IPO Shares,Listed Shares,Equity,Warrant,Nominal,MCB,Is Affiliated,Is Acceleration,Is New,Lock Up,Subscribed Stock,All UW,Amount\n"
	helpers.ExportResponse(filePath, header, listStock)
	fmt.Println("Finish Export By Underwriter")
}

func (service *StockServiceImpl) ExportByValue(value int, underwriter string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	listStock, err := service.Stock_Repository.FindByValue(ctx, value, underwriter)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(listStock) == 0 {
		fmt.Println("No stock found")
		return
	}

	filePath := fmt.Sprintf("output/value/value_%s_%d.csv", underwriter, value)
	header := "Stock Code,price,IPO Shares,Listed Shares,Equity,Warrant,Nominal,MCB,Is Affiliated,Is Acceleration,Is New,Lock Up,Subscribed Stock,All UW,Amount\n"
	helpers.ExportResponse(filePath, header, listStock)
	fmt.Printf("Finish Export Value %s underwriter\n", underwriter)
}
