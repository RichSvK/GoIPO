package services

import (
	"IPO/helpers"
	"IPO/repository"
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

type IpoDetailService interface {
	InsertDetail(fileName string)
}

type IpoDetailServiceImpl struct {
	IpoDetailRepository repository.IpoDetailRepository
}

func NewDetailService(detailRepository repository.IpoDetailRepository) IpoDetailService {
	return &IpoDetailServiceImpl{
		IpoDetailRepository: detailRepository,
	}
}

func (service *IpoDetailServiceImpl) InsertDetail(fileName string) {
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
			fmt.Println("Error opening file:", err)
		}
	}()

	reader := bufio.NewReader(file)
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	// read & ignore header
	if _, err := csvReader.Read(); err != nil {
		if err == io.EOF {
			return
		}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("csv read error:", err)
			return
		}

		detail := helpers.SplitDetailString(record)

		if err := service.IpoDetailRepository.Save(ctx, detail); err != nil {
			fmt.Println("Error insert IPO detail:", err)
			return
		}
	}

	fmt.Println("Success Insert IPO Detail Data")
}
