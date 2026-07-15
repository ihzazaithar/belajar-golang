package main

import (
	"fmt"
)

func main() {

	var varArr = [10]int{1, 2, 3, 4}

	fmt.Println(varArr[1])

	for k, v := range varArr {

		fmt.Println("Index -", k, "Is ", v)
	}

	varSlices := varArr[:2]

	fmt.Println(varSlices[0])

}
