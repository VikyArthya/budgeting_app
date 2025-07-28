package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(){
	// Connection string untuk PostgreSQL
    dsn := "host=localhost user=viky_arthya password=Vikysan0 dbname=budgeting port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error()) // Tambahkan err.Error() untuk detail
    }

    // Tes koneksi
    sqlDB, err := db.DB()
    if err != nil {
        panic("failed to get database instance: " + err.Error())
    }
    if err := sqlDB.Ping(); err != nil {
        panic("failed to ping database: " + err.Error())
    }
    log.Println("Successfully connected to database!")	
}