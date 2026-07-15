package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	var varRing *ring.Ring = ring.New(5) // Perlu dimasukan banyaknya data

	// Input cara manual

	// varRing.Value = "Data 1"

	// varRing = varRing.Next()
	// varRing.Value = "Data 2"

	// varRing = varRing.Next()
	// varRing.Value = "Data 3"

	// varRing = varRing.Next()
	// varRing.Value = "Data 4"

	// varRing = varRing.Next()
	// varRing.Value = "Data 5"

	// Input dengan for
	for e := 0; e < varRing.Len(); e++ {
		varRing.Value = "Data " + strconv.Itoa(e)
		varRing = varRing.Next()
	}

	// Memanggil nilai pada struktur data ring

	// fmt.Println(varRing.Value)

	// memanggil dengan fungsi

	varRing.Do(func(value any) {
		fmt.Println(value)
	})

}
