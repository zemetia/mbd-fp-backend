package config

import (
	"fmt"
	"fp-mbd-amidrive/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entity.User{},
		entity.Lokasi{},
		entity.Membership{},
		entity.MetodePembayaran{},
		entity.Mobil{},
		entity.ReviewMobil{},
		entity.TipeMesin{},
		entity.TipeMobil{},
		entity.TipePersneling{},
		entity.Transaksi{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}
