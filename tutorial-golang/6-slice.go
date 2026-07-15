package main

import "fmt"

func main() {

	var varArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(varArray)

	// Opsi 1
	varSlices1 := varArray[2:6] // varArray[firstRange:lastRange]
	// Memotong dari index ke 2 dan berakhir di index ke 6. Index terakhir memang 6, tetapi yang ditampilkan adalah index ke 5nya.
	// secara penulisan [2:6] tetapi yang sebenarnya ditampilkan adalah [2:6-1].
	fmt.Println(varSlices1)

	// Opsi 2
	varSlices2 := varArray[2:] // varArray[firstRange:]
	// Memotong dari index ke X hingga index terakhir dari array tersebut.
	fmt.Println(varSlices2)

	// Opsi 3
	varSlice3 := varArray[:6] // varArray[:lastRange]
	// Memotong dari index pertama hingga index ke X yang inginkan.
	fmt.Println(varSlice3)

	// Opsi 4
	varSlice4 := varArray[:]
	// Memotong array dari index pertama hingga terakhir atau semua data Arraynya diambil.
	fmt.Println(varSlice4)

	fmt.Println("==================================")
	// Perilaku unik slice/array. mengubah data slice yang mengambil dari suatu array, maka nilai arraynya juga akan berubah.
	var varSlice5 = varArray[4:6]
	varSlice5[0] = 1
	varSlice5[1] = 2

	fmt.Println(varArray)
	fmt.Println(varSlice5)

	fmt.Println("=============Fungsi===============")

	varSlice6 := append(varSlice5, 11) // Menambahkan data, jika slice sudah penuh dari capacitynya. maka akan dibuatkan array baru.
	varSlice7 := make([]int, 5, 10)    // Membuat slices kosong

	fmt.Println(cap(varSlice5)) // Melihat jumlah capacity dari suatu slice.
	fmt.Println(varSlice6)
	fmt.Println(varSlice7)
}
