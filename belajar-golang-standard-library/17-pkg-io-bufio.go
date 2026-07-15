package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Read
	input := strings.NewReader("this is long string\nwith new line\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	// Write

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello World\n")           // Jika Ingin menangkap return errornya juga
	varInt, _ := writer.WriteString("Selamat Belajar\n") // Jika ingin menangkap nilai Int dan returnya
	writer.WriteString("Semoga Sukses!\n")               // Jika tidak ingin menangkap data
	writer.Flush()                                       // Setiap selesai menulis
	fmt.Println(varInt)
}
