package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Base64 Encode/Decode
	value := "Ihza Zaidan Afthar"

	encoded := base64.StdEncoding.EncodeToString([]byte(value)) // harus slice []byte
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	// Karena nilai mengembalikan error yang disebabkan mungkin karena tipe data yang salah, dicheck terlebih dahulu.
	if err != nil { // err == nill berarti tidak error, kalau err != nill artinya error
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(decoded)
		fmt.Println(string(decoded))
	}

	// CSV Reader
	csvString := "ihza, zaidan, afthar\n" + "budi, pratama, nugraha\n" + "joko, morro, diah"
	// fmt.Println(csvString)

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	// CSV Writer
	writer := csv.NewWriter(os.Stdout) // Menuliskan ke layar; menuliskannya bisa lain tempat seperti file, dll.

	_ = writer.Write([]string{"ihza", "zaidan", "aftar"})
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush() // Gunakan ini setelah selesai menulis
}
