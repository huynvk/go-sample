package main

import (
	"fmt"
	"os"
)

func write(fileName string, content string) {
	f, e := os.Create(fileName)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	fmt.Fprintln(f, content)
}

func increase() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	write("test.txt", "Sample content")
	fmt.Println(increase())
}
