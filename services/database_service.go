package services

import (
	"IPO/configs"
	"IPO/models/entity"
	"fmt"
)

func CreateDatabaseTable() {
	err := configs.PoolDB.AutoMigrate(&entity.Stock{}, &entity.Broker{}, &entity.IPO_Detail{})
	if err != nil {
		fmt.Println("Error Migrating")
		return
	}
	fmt.Println("Success migrating")
}

func DeleteTable() {
	db := configs.PoolDB

	tables := []interface{}{
		&entity.IPO_Detail{},
		&entity.Stock{},
		&entity.Broker{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			fmt.Printf("Failed to drop %T: %v\n", table, err)
			return
		}
	}

	fmt.Println("Success Drop Table")
}

func ClearTable() {
	TruncateTable("ipo_detail")
	TruncateTable("stock_ipo")
	TruncateTable("broker_underwriter")
	fmt.Println("Success Truncate Table")
}

func TruncateTable(tableName string) {
	db := configs.PoolDB
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	defer func() {
		db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	}()

	result := db.Exec("TRUNCATE TABLE " + tableName)
	if result.Error != nil {
		fmt.Println("Error Truncate Table")
		return
	}
}
