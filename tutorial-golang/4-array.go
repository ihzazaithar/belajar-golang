package main

import "fmt"

func main() {

	var varArrayKosong [5]int

	fmt.Println(varArrayKosong)

	var varArrayBerisiFixLength1 = [5]int{1, 2, 3, 4}
	var varArrayBerisiFixLength2 = [5]int{1, 2, 3, 4, 5}
	// Panjangnya 5, tetapi bisa hanya diisi 5 atau kurang. yang tidak terisi maka akan diisi oleh nilai defaultnya

	fmt.Println(varArrayBerisiFixLength1)
	fmt.Println(varArrayBerisiFixLength2)

	var varArrayBerisiNoFixLength = [...]int{1, 2, 3} // Max panjang arraynya mengikuti jumlah data yang diisi pertama kali
	fmt.Println(varArrayBerisiNoFixLength)

}
