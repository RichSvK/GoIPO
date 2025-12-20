package services

import (
	"IPO/helpers"
	"IPO/models/entity"
	"IPO/repository"
	"bufio"
	"context"
	"encoding/csv"
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
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	// Read & ignore header
	if _, err := csvReader.Read(); err != nil {
		if err == io.EOF {
			fmt.Println("File is empty")
			return
		}
		fmt.Println("Error reading header:", err)
		return
	}

	const batchSize = 1000
	stocks := make([]entity.Stock, 0, batchSize)
	lineNumber := 1
	successCount := 0
	errorCount := 0

	// Create a context with timeout for database operations with 30 seconds limit since batch inserts can take longer
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			// Insert remaining records
			if len(stocks) > 0 {
				if err := service.Stock_Repository.SaveBatch(ctx, stocks); err != nil {
					fmt.Printf("Error inserting final batch: %v\n", err)
					errorCount += len(stocks)
				} else {
					successCount += len(stocks)
				}
			}
			break
		}
		if err != nil {
			fmt.Printf("CSV read error at line %d: %v\n", lineNumber+1, err)
			errorCount++
			lineNumber++
			continue
		}
		lineNumber++

		// Parse the stock data
		stock := helpers.SplitStockString(record)
		stocks = append(stocks, stock)

		// Insert batch of stocks when it reaches batchSize
		if len(stocks) >= batchSize {
			if err := service.Stock_Repository.SaveBatch(ctx, stocks); err != nil {
				fmt.Printf("Error inserting batch at line %d: %v\n", lineNumber, err)
				errorCount += len(stocks)
			} else {
				successCount += len(stocks)
			}
			stocks = stocks[:0]
		}
	}

	fmt.Printf("Insert completed: %d successful, %d errors\n", successCount, errorCount)
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
