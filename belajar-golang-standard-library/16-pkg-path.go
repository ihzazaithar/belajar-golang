package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {

	// Path
	fmt.Println(path.Dir("hello/world.go"))             // Membaca dan menampilkan yang merupakan direktori dari path tersebut -- hello adalah direktori
	fmt.Println(path.Base("hello/world.go"))            // Membaca dan menampilkan yang merupakan file dari path tersebtu --> world.go adalah file
	fmt.Println(path.Ext("hello/world.go"))             // Membaca dan menampilkan yang merupakan ekstensi file dari path tersebut --> .go adalah ekstensi
	fmt.Println(path.Join("hello", "world", "main.go")) // Menggabungkan string untuk dijadikan url --> hello/world/main.go

	fmt.Println("=======================================")
	// path/filepath --> digunakan untuk file system seperti Windows yang menggunakan backslash \
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go")) // Untuk mengetahui apakah path tersebut merupakan Relative Path atau bukan
	fmt.Println(filepath.IsAbs("hello/world.go"))   // Untuk mengetahui apakah path tersebut merupakan Absolute Path atau bukan.
	// Absolute Path adalah path berisikan dari rootnya
	// dan, relative path adalah path berisikan dari posisi kita berada.
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
