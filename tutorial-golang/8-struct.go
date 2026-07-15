package main

import "fmt"

// Opsi 1 Deklarasi Struct
type Negara struct {
	Nama     string
	Ibukota  string
	Provinsi []string
}

func main() {
	// Opsi 1 Pemanggilan
	var Negara1 Negara

	Negara1.Nama = "Indonesia"
	Negara1.Ibukota = "Jakarta"
	Negara1.Provinsi = []string{"Jawa Barat", "Jakarta"}

	fmt.Println(Negara1)

	// Opsi 2 Pemanggilan (Struct Literal)
	Negara2 := Negara{
		Nama:     "Malaysia",
		Ibukota:  "Kuala Lumpur",
		Provinsi: []string{"KL1", "KL2"},
	}

	fmt.Println(Negara2)

	// Opsi 3 Pemanggilan
	// Opsi ini pastikan bahwa urutan input dan urutan deklarasi struct sama.
	Negara3 := Negara{"Thailand", "Bangkok", []string{"Pattaya", "Phuket"}}
	fmt.Println(Negara3)

}
