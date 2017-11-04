package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPUs/Cores:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(runtime.NumCPU()))
}
