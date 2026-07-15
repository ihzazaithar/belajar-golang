package main

import (
	"container/list"
	"fmt"
)

func main() {

	var data *list.List = list.New()

	// Input data ke list ke urutan awal
	data.PushBack("Ihza")
	data.PushBack("Zaidan")
	data.PushBack("Afthar")
	// atau menginput ke urutan akhir atau ke posisi N dengan cara lain.
	// Akses value list
	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	// atau dengan for loop

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
