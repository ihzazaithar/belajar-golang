package main

import "fmt"

func main() {

	misal := getGoodBye
	contoh := getGoodBye("Ihza")

	fmt.Println(misal("Zaidan"))
	fmt.Println(contoh)

}

func getGoodBye(name string) string {

	return "Good Bye " + name

}
