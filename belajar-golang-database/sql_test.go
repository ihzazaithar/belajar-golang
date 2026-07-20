package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Query Data tanpa Mengembalikan Result (INSERT, DELETE, UPDATE)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES('ihza', 'Ihza')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}

// Query Data dengan Mengembalikan Result (SELECT)

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // Berjalan jika rows.Next() mengembalikan true
		var id, name string         // Samakan tipe datanya dengan tipe data yang ada di database.
		err = rows.Scan(&id, &name) // Posisi urutannya berdasarkan dari query SELECTnya. Gunakan pointer
		if err != nil {
			panic(err)
		}
		fmt.Println("Id", id)
		fmt.Println("Name", name)
	}

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married, create_at FROM customer" // Deklarasikan column satu per satu, jangan gunakan * asterisk.
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // Berjalan jika rows.Next() mengembalikan true
		var id, name string      // Samakan tipe datanya dengan tipe data yang ada di database.
		var email sql.NullString // Tipe data ini untuk menangani data NULL
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime // Tipe data ini menangani data NULL
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) // Posisi urutannya berdasarkan dari query SELECTnya. Gunakan pointer
		if err != nil {
			panic(err)
		}
		fmt.Println("==================================")
		fmt.Println("Id", id)
		fmt.Println("Name", name)
		// fmt.Println("Email", email)
		if email.Valid {
			fmt.Println("Email", email.String)
		}
		fmt.Println("Balance", balance)
		fmt.Println("Rating", rating)
		// fmt.Println("Birth Date", birthDate)
		if birthDate.Valid {
			fmt.Println("Birth Date", birthDate.Time)
		}
		fmt.Println("Married", married)
		fmt.Println("Created At", createdAt)

	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// username := "admin'; #" // Query SQL Injection
	username := "admin"
	password := "admin"

	ctx := context.Background()
	// Jangan gunakan ini dalam implementasi aslinya
	script := "SELECT username FROM user WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

// ** Solusi SQL Injection

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// username := "admin'; #" // Query SQL Injection
	username := "admin"
	password := "admin"

	ctx := context.Background()
	script := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "joko"
	password := "joko"

	ctx := context.Background()
	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}

// Solusi SQL Injection *

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "joko@gmail.com"
	comment := "Test Komen"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	ctx := context.Background()
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}

// ** Prepare Statment

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script) // Binding Prepare statement dan koneksi database
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "joko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id ", id)
	}

}

// Prepare Statement *

// ** Database Transaction
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// do any transaction
	for i := 0; i < 10; i++ {
		email := "joko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id ", id)
	}

	// err = tx.Commit()   // Komit ke Database
	err = tx.Rollback() // Membatalkan Proses Transaction
	if err != nil {
		panic(err)
	}

}
