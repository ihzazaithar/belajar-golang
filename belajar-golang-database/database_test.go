package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Memanggil function init() dan _ artinya agar tidak perlu dipakai
)

func TestEmpty(t *testing.T) {

}

// ** Membuat Koneksi / Database Pooling

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "Izatar:@tcp(localhost:3306)/belajar_golang_database") // Validasi Format Argument benar, belum koneksi fisik.

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Gunakan Database

}
