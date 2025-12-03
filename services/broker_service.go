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

type BrokerService interface {
	InsertBroker(fileName string)
}

type BrokerServiceImpl struct {
	BrokerRepository repository.BrokerRepository
}

func NewBrokerService(brokerRepository repository.BrokerRepository) BrokerService {
	return &BrokerServiceImpl{
		BrokerRepository: brokerRepository,
	}
}

func (service *BrokerServiceImpl) InsertBroker(fileName string) {
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

		broker := helpers.SplitBrokerString(result)
		err = service.BrokerRepository.Save(ctx, broker)
		if err != nil {
			fmt.Println("Error inserting broker data")
			return
		}
	}
	fmt.Println("Success Insert Broker Data")
}
