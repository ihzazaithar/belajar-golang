package main

import "fmt"

// Opsi 1 Deklarasi Struct
type Negara struct {
	Nama     string
	Ibukota  string
	Provinsi []string
}

func (x Negara) main1(y string) string {

	return y + " " + x.Nama

}

func main() {

	// Opsi 1 Pemanggilan
	var Negara1 Negara

	Negara1.Nama = "Indonesia"
	Negara1.Ibukota = "Jakarta"
	Negara1.Provinsi = []string{"Jawa Barat", "Jakarta"}

	z := Negara1.main1("Hello")

	fmt.Println(z)

}
