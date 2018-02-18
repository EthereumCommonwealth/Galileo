package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"log"
	"github.com/EthereumCommonwealth/Galileo/models"
)

func getDBConnection() {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		log.Fatal(fmt.Printf("DB connection error %v", err))
	}
	defer db.Close()
}

func migrateBlockTable(db *gorm.DB) {
	db.AutoMigrate(&models.Block{})

	// Custom block migration here
}