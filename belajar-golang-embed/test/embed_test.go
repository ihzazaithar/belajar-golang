package test

import (
	"embed"
	_ "embed" // Mengimport tetapi tidak digunakan perlu ditambahkan _
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string // File embed harus selalu di luar Function

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.jpg
var logo []byte // File embed harus selalu di luar Function, tipe data ini digunakan jika file berisikan binary

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_new.jpg", logo, fs.ModePerm) // Membuat/Mengcopy variable embed ke file varu
	if err != nil {
		panic(err)
	}
}

// Manual Multiple Files

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

// Path Match Multiple Files

//go:embed files/*.txt
var path embed.FS

func TestPatchMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			fileContent, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(fileContent))
		}
	}
}
