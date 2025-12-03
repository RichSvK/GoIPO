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

	// Remove header
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

		detail := helpers.SplitDetailString(result)
		err = service.IpoDetailRepository.Save(ctx, detail)
		if err != nil {
			fmt.Println("Error insert IPO detail")
			return
		}
	}
	fmt.Println("Success Insert IPO Detail Data")
}
