package utilities

import (
	"IPO/helpers"
	"fmt"
	"strconv"
)

func MainMenu() int {
	choiceString := ""
	var err error = nil
	choice := 0

	for {
		helpers.ClearScreen()
		fmt.Println("Main Menu")
		fmt.Println("1. Export Data")
		fmt.Println("2. Insert Data")
		fmt.Println("3. Database Menu")
		fmt.Println("4. Exit")
		choiceString = helpers.ScanInput(">> ")
		choice, err = strconv.Atoi(choiceString)
		if err != nil || (choice < 1 || choice > 4) {
			fmt.Println("Invalid Input")
			helpers.PressEnter()
		} else {
			break
		}
	}

	return choice
}
