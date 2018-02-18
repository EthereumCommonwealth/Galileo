package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"log"
	"github.com/EthereumCommonwealth/Galileo/common"
)

func GetDBConnection(setting common.GalileoSetting) *gorm.DB {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		setting.DBHost, setting.DBPort, setting.DBUser, setting.DBName, setting.DBPass)
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(fmt.Printf("DB connection error %v", err))
	}
	return db
}

func Migrate(db *gorm.DB) {
	migrateBlockTable(db)
	migrateAddressTable(db)
	migrateTransactionTable(db)
}

func migrateBlockTable(db *gorm.DB) {
	db.AutoMigrate(Block{})

	// Custom block migration here
}

func migrateAddressTable(db *gorm.DB) {
	db.AutoMigrate(Address{})

	// Custom address migration here
}

func migrateTransactionTable(db *gorm.DB) {
	db.AutoMigrate(Transaction{})

	// Custom Transaction migration here
}