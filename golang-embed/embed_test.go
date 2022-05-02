package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println("version", version)
}

// go:embed logo.png
var logo []byte

func TestByteArrayCreateNewFile(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestByteArrayReadFile(t *testing.T) {
	file, err := ioutil.ReadFile("logo.png")
	if err != nil {
		panic(err)
	}

	fmt.Println("file", file)
}

//go:embed multiple-files/a.txt
//go:embed multiple-files/b.txt
//go:embed multiple-files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("multiple-files/a.txt")
	fmt.Println("a", string(a))

	b, _ := files.ReadFile("multiple-files/b.txt")
	fmt.Println("b", string(b))

	c, _ := files.ReadFile("multiple-files/c.txt")
	fmt.Println("c", string(c))
}

//go:embed multiple-files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("multiple-files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("fileName", entry.Name())
			content, _ := path.ReadFile("multiple-files/" + entry.Name())
			fmt.Println("content", string(content))
		}
	}
}
