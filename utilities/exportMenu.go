package utilities

import (
	"IPO/helpers"
	"IPO/repository"
	"IPO/services"
	"fmt"
	"strconv"
)

func ExportMenu() {
	userChoice := ""
	choice := 0
	var err error = nil
	for {
		helpers.ClearScreen()
		fmt.Println("Export Menu")
		fmt.Println("1. Export By Underwriter")
		fmt.Println("2. Export By Value Amount")
		fmt.Println("3. Return to Main Menu")
		userChoice = helpers.ScanInput("Choice [1 - 3]: ")
		choice, err = strconv.Atoi(userChoice)
		if err == nil && (choice > 0 && choice < 4) {
			break
		}
		fmt.Println("Invalid Input")
		helpers.PressEnter()
	}

	switch choice {
	case 1:
		ExportByUW_Menu()
	case 2:
		ExportByAmmount_Menu()
	}
}

func ExportByUW_Menu() {
	helpers.ClearScreen()
	fmt.Println("Export By Underwriter Menu")

	var underwriter string = ""
	var inputLength int = 0
	valid := true
	for {
		underwriter = helpers.ScanInput("Type UW Code [2 Character and Uppercase][Type 0 to Return]: ")
		inputLength = len(underwriter)

		if inputLength != 2 {
			if underwriter == "0" {
				return
			}
			fmt.Println("UW Code must be 2 character")
			valid = false
		} else {
			valid = true
			for i := 0; i < inputLength; i++ {
				if underwriter[i] < 'A' || underwriter[i] > 'Z' {
					valid = false
					break
				}
			}
		}

		if valid {
			break
		}
	}

	stockService := services.NewStockService(repository.NewStockRepository())
	stockService.ExportByUnderwriter(underwriter)
}

func ExportByAmmount_Menu() {
	var valueChoice string = ""
	var value int = 0
	var err error = nil
	for {
		helpers.ClearScreen()
		fmt.Println("Export By Ammount Menu")
		fmt.Println("1. Group 1 (Value <= 250 Billion)")
		fmt.Println("2. Group 2 (Value 250 - 500 Billion)")
		fmt.Println("3. Group 3 (Value 500 Billion - 1 Trillion)")
		fmt.Println("4. Group 4 (Value > 1 Trillion)")
		fmt.Println("5. Return to Main Menu")
		valueChoice = helpers.ScanInput("Choice[1 - 5]: ")
		value, err = strconv.Atoi(valueChoice)
		if err == nil && (value > 0 && value < 6) {
			break
		}
		fmt.Println("Invalid")
		helpers.PressEnter()
	}

	if value == 5 {
		return
	}

	underwriter := ""
	valid := false
	stockService := services.NewStockService(repository.NewStockRepository())
	for {
		underwriter = helpers.ScanInput("Input Underwriter [ALL | 2 Uppercase Letter]: ")

		if len(underwriter) == 2 {
			valid = true
			for i := 0; i < 2; i++ {
				if underwriter[i] < 'A' || underwriter[i] > 'Z' {
					valid = false
					break
				}
			}
		}

		if valid || underwriter == "ALL" {
			break
		}

		fmt.Println("Invalid")
		helpers.PressEnter()
	}
	stockService.ExportByValue(value, underwriter)
}
