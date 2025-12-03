package main

import (
	"IPO/configs"
	"IPO/helpers"
	"IPO/utilities"
	"fmt"
)

func init() {
	configs.MakeOutputFolder("output")
	configs.InitEnvironment()
	configs.OpenConnection()
}

func main() {
	defer func() {
		if err := configs.SqlDB.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	choice := 0
	for {
		helpers.ClearScreen()
		choice = utilities.MainMenu()
		switch choice {
		case 1:
			utilities.ExportMenu()
		case 2:
			utilities.InsertMenu()
		case 3:
			utilities.DatabaseMenu()
		}

		if choice == 4 {
			break
		}
		helpers.PressEnter()
	}
	fmt.Println("Program Finished")
}
