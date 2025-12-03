package utilities

import (
	"IPO/helpers"
	"IPO/services"
	"fmt"
	"strconv"
)

func DatabaseMenu() {
	userChoice := ""
	choice := 0
	var err error
	for {
		helpers.ClearScreen()
		fmt.Println("Database Menu")
		fmt.Println("1. Create Database Table")
		fmt.Println("2. Clear Table Data")
		fmt.Println("3. Delete Database Table")
		fmt.Println("4. Return to Main Menu")
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
		services.CreateDatabaseTable()
	case 2:
		services.ClearTable()
	case 3:
		services.DeleteTable()
	case 4:
		return
	}
}
