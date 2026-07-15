package main

import "fmt"

func main() {

	// Opsi 1
	varMap1 := map[string]string{"Nama": "Kaiju"} // Langsung mendeklarasikan Datanya

	// Opsi 2
	varMap2 := map[string]string{} // Tanpa Mendeklarasikan Data

	fmt.Println(varMap1["Nama"])
	fmt.Println(varMap2)

	fmt.Println("==============Fungsi==================")

	varMap2["Nama"] = "Ihza"
	varMap2["Kota"] = "Depok" // Mengganti atau Menambahkan Value terhadap Keynya
	fmt.Println(len(varMap2))
	fmt.Println(varMap2["Nama"]) // Mengakses Value dengan menggunakan Key di Map

	varMap3 := make(map[string]string, 100) // Opsi 3 untuk membuat map. perbedaan dengan opsi kedua adalah opsi ini menentukan panjang dari data Mapnya. opsi 1, opsi 2, dan opsi 3, akan menambahkan panjangnya ketika defaultnya/panjang yang ditentukan sudah tidak cukup.
	fmt.Println(varMap3)

}
