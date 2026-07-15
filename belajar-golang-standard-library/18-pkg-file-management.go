package main

import (
	"bufio"
	"io"
	"os"
)

// Fungsi untuk membuat file baru dan menuliskan ke sana.
func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func bacaFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func tambahKonten(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createNewFile("sample.txt", "This is sampel text")

	// result, _ := bacaFile("sample.txt")
	// fmt.Println(result)

	tambahKonten("sample.txt", "\nTambahan Data")

}
