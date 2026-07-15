package main

import "fmt"

func main() {

	var varArray [7]string
	fmt.Println(varArray)

	varArray[0] = "Senin"
	varArray[1] = "Selasa"
	varArray[2] = "Rabu"
	varArray[3] = "Kamis"
	varArray[4] = "Jum'at"
	varArray[5] = "Sabtu"
	varArray[6] = "Minggu"
	// Mengganti Nilai array dengan Indexnya.
	fmt.Println(varArray)

	varArray[1] = "Selasa-Libur" // Mengganti nilai array dari Index
	fmt.Println(varArray)
	fmt.Println("Index ke 2 dari varArray adalah: ", varArray[2]) // Menampilkan/Mengambil nilai array di posisi Indexnya
	fmt.Println("Panjang varArray: ", len(varArray))              // Menampilkan Jumlah/Panjang Array

}
