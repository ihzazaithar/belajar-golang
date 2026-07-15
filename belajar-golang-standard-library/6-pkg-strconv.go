package main

import (
	"fmt"
	"strconv"
)

func main() {

	result, err := strconv.ParseBool("saaal")
	// strvonc.X() --> urutan pengembaliannya adalah value tipe datanya dan error.
	// saat konversi bisa terjadi error, makanya perlu untuk dilakukan pengecekan.
	if err == nil {
		fmt.Println(result)
	} else if err != nil {
		fmt.Println(err.Error())
	}

}
