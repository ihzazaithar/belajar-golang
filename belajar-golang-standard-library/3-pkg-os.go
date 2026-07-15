package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args akan menangkap argumen saat menjalankan go run namaFile args1 args2 argsN dalam bentuk slices.
	// Argumen dipisahkan oleh spasi, untuk mengirim satu argumen yang memiliki spasi gunakan ""
	args := os.Args

	for _, data := range args { // _ --> blank, posisi menyimpan index. data --> posisi menyimpan valuenya
		fmt.Println(data)
	}

	hostnm, err := os.Hostname() // Mengembalikan nilai dengan urutan: string (nama hostname), error

	if err == nil {
		fmt.Println(hostnm)
	} else if err != nil {
		fmt.Println("Hostname tidak ada", err.Error())
	}

}
