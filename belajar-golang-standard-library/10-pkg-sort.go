package main

import (
	"fmt"
	"sort"
)

type User struct {
	nama string
	umur int
}

type UserSlice []User // Interface

// Kontrak yang akan dipakai untuk package sort
func (a UserSlice) Len() int {
	return len(a)
}

func (a UserSlice) Less(i, j int) bool {
	return a[i].umur < a[j].umur
}

func (a UserSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i] // Menukar value antar variable
}

func main() {

	users := []User{
		{"Eko", 30},
		{"Budi", 25},
		{"Joko", 35},
		{"Adit", 20},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)

}
