package configs

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var JwtSecret = "yoursecretkey"
var DB *gorm.DB

func ConnectDatabase() {
	// Mulai waktu koneksi
	startTime := time.Now()

	// Buat context dengan timeout 5 detik
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := "root:@tcp(localhost:3306)/restAPI?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("Connecting to database...")

	// Koneksi ke database dengan context
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Cek koneksi dengan Ping
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Gunakan context untuk ping
	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	// Catat waktu koneksi selesai
	duration := time.Since(startTime)
	log.Printf("Database connected successfully in %v\n", duration)

	// Simpan koneksi database ke variabel global
	DB = database
}
