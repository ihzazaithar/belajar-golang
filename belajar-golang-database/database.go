package belajar_golang_database

import (
	"database/sql"
	"time"
)

// Membuat Function Koneksi agar tidak perlu deklarasi terus menerus.
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true") // Validasi Format Argument benar, belum koneksi fisik. parseTime=true diperlukan untuk menangani tipe data date, datetime, timestamp secara otomatis diparsing di Golang.

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
