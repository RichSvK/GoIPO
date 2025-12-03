package utilities

import (
	"IPO/helpers"
	"IPO/repository"
	"IPO/services"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func InsertMenu() {
	userChoice := ""
	choice := 0
	var err error
	for {
		helpers.ClearScreen()
		fmt.Println("Insert Menu")
		fmt.Println("1. Insert Data IPO Stock")
		fmt.Println("2. Insert Data Underwriter")
		fmt.Println("3. Insert Data IPO Detail")
		fmt.Println("4. Exit")
		userChoice = helpers.ScanInput("Choice [1 - 4]: ")
		choice, err = strconv.Atoi(userChoice)
		if err != nil || (choice < 1 || choice > 4) {
			fmt.Println("Invalid Input")
			helpers.PressEnter()
		} else {
			break
		}
	}

	switch choice {
	case 1:
		var path = filepath.Join("data", "stock")
		stockService := services.NewStockService(repository.NewStockRepository())
		stockService.InsertStock(GetInsertFilePath(path, "Menu Insert IPO Stock"))
	case 2:
		var path = filepath.Join("data", "broker")
		brokerService := services.NewBrokerService(repository.NewBrokerRepository())
		brokerService.InsertBroker(GetInsertFilePath(path, "Menu Insert Underwriter"))
	case 3:
		var path = filepath.Join("data", "detail")
		detailService := services.NewDetailService(repository.NewDetailRepository())
		detailService.InsertDetail(GetInsertFilePath(path, "Menu Insert IPO Detail"))
	case 4:
		return
	}
}

func GetInsertFilePath(folderName string, menu string) string {
	helpers.ClearScreen()
	fmt.Println(menu)

	fileList, err := helpers.ReadFolder(folderName)

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	size := len(fileList)
	for i := 0; i < size; i++ {
		tempName := strings.Split(fileList[i], folderName)
		fmt.Printf("%d. %s from %s\n", (i + 1), tempName[1], fileList[i])
	}

	var choice int
	var userInput string
	promptString := fmt.Sprintf("Input [1 - %d] [Type 0 to return]: ", size)
	for {
		userInput = helpers.ScanInput(promptString)
		choice, err = strconv.Atoi(userInput)
		if err == nil && choice >= 0 && choice <= size {
			break
		} else {
			fmt.Println("Invalid Input")
		}
	}

	if choice == 0 {
		return ""
	}
	return fileList[choice-1]
}
