package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.TipeData(Variable, DefaultValue, Description)
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "password", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 553, "host port")

	flag.Parse()

	fmt.Println("Username:", *host)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *username)
	fmt.Println("Port:", *port)
}
