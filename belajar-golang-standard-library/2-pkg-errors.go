package main

import (
	"errors"
	"fmt"
)

// Best Practice Deklarasi Error Custom Message Programmer Golang
var (
	ValidationError = errors.New("Validasi gagal!")
	NotFoundError   = errors.New("Data tidak ada!")
)

func getById(data string) error {

	if data == "" {
		return ValidationError
	} else if data != "kucing" {
		return NotFoundError
	}

	// Jika data aman, tampilkan
	return nil

}

func main() {

	err := getById("kucing")

	// Menghandle/Menampilkan pesan error yang ditangkap.
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown Errors")
		}
	} else {
		fmt.Println("Data Aman!")
	}

}
