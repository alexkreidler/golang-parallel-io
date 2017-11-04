package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	wg.Add(1)
	lsFiles(dir)
	wg.Wait()
}

var wg sync.WaitGroup

var (
	concurrent    = runtime.GOMAXPROCS(runtime.NumCPU())
	semaphoreChan = make(chan struct{}, concurrent)
)

func lsFiles(dir string) {
	// block while full
	semaphoreChan <- struct{}{}

	go func() {
		defer func() {
			// read to release a slot
			<-semaphoreChan
			wg.Done()
		}()

		file, err := os.Open(dir)
		if err != nil {
			fmt.Println("error opening directory", err.Error())
		}
		defer file.Close()
		files, err := file.Readdir(-1) // Loads all children files into memory. A more efficient way?
		if err != nil {
			fmt.Println("error reading directory", err.Error())
		}
		for _, f := range files {
			fmt.Println(dir + "/" + f.Name())
			if f.IsDir() {
				wg.Add(1)
				go lsFiles(dir + "/" + f.Name())
			}
		}
	}()
}
