package main

import (
        "fmt"
        "log"
        "os"
        "path/filepath"
)

func main() {
        err := filepath.Walk(os.Args[1], func(path string, fi os.FileInfo, err error) error {
                if err != nil {
                        return err
                }
                fmt.Println(path)
                return nil
        })
        if err != nil {
                log.Fatal(err)
        }
}
