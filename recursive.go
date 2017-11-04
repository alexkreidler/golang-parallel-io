package main

import (
	"fmt"
	"os"
)

func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	lsFiles(dir)
}

func lsFiles(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("error opening directory")
	}
	defer file.Close()
	files, err := file.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory")
	}
	for _, f := range files {
		if f.IsDir() {
			lsFiles(dir + "/" + f.Name())
		}
		fmt.Println(dir + "/" + f.Name())
	}
}
