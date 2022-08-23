package database

import (
	"fmt"
	"golang/models"
	"golang/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Product{},
		&models.Toping{},
		&models.Transaction{},
		&models.Cart{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
