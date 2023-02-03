package connection

import (
	"fmt"
	"log"
	"os"
	"uts/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatalln(" Error to Load Env")
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't Connect To Database")
	}

	db.AutoMigrate(&models.Peminjaman{})
	db.AutoMigrate(&models.Petugas{})
	db.AutoMigrate(&models.Pengunjung{})
	db.AutoMigrate(&models.Buku{})

	return db
}

