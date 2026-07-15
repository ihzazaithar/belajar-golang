package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Ihza Zaidan", "Ihza"))                 // Mencari nilai "Ihza" dari strings "Ihza Zaidan"
	fmt.Println(strings.Split("Ihza-Zaidan", "-"))                       // Memisahkan suatu string berdasarkan tanda -. tanda disini bisa macam-macam, dari spasi, -, *, dll
	fmt.Println(strings.ToLower("Ihza Zaidan"))                          // Mengubah string ke huruf kecil semua
	fmt.Println(strings.ToUpper("Ihza Zaidan"))                          // Mengubah string ke huruf besar semua
	fmt.Println(strings.Trim("-----Ihza Zaidan------", "-"))             // Menghapus semua - yang ada di belakang atau depan. tanda - bisa diubah sesuai yang ingin dihapus.
	fmt.Println(strings.ReplaceAll("Ihza Zaidan Afthar", "Afthar", "A")) // Mengganti suatu yang ada di string menjadi yang lain. di sini saya mengubah Afthar menjadi A.

}
