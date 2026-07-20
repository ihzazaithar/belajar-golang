package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed logo.jpg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			// fileContent, err :=
			fileContent, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(fileContent))
		}
	}
}
