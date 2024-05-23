package configs

import (
	"github.com/joho/godotenv"
)

func init() {
	InitEnvironment()
	OpenConnection()
	MakeOutputFolder("output")
	MakeOutputFolder("output/value")
	MakeOutputFolder("output/UW")
}

func InitEnvironment() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}
}
