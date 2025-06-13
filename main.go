package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cr92/gosht/dataSrc"
	"github.com/cr92/gosht/file"
)

func main() {
	start := time.Now()

	loadEnv()

	dataFile := file.File{Path: os.Getenv("DATA")}
	processData(&dataFile)

	fmt.Println(time.Since(start))
}

func processData(ds dataSrc.DataSrc) {
	dest := make(chan string)
	defer close(dest)

	done := make(chan bool)
	defer close(done)

	go ds.ReadLine(dest, done)

	for {
		select {
		case m := <-dest:
			fmt.Println(m)
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}
