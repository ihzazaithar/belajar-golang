package main

import (
	"fmt"
	// "go_say_hello" merupakan nama package dari module yang kita import, bisa diubah untuk memudahkan atau kenyamanan.
	go_say_hello "github.com/ihzazaithar/go-say-hello/v2"
)

func main() {
	fmt.Println(go_say_hello.SayHello("Ihza"))
}
