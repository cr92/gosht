package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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
	var wg sync.WaitGroup

	dest := make(chan string)
	wg.Add(1)
	go func() {
		ds.ReadLine(dest)
		close(dest)
		wg.Done()
	}()

	for i := range 10 {
		wg.Add(1)
		go func(dest chan string, id int) {
			count := 0
			f, err := os.Create(fmt.Sprintf("batch_%d.csv", id))
			if err != nil {
				log.Fatal("Error creating file")
			}
			defer func() {
				err := f.Close()
				if err != nil {
					log.Fatal("Error closing file")
				}
			}()
			defer wg.Done()

			for {
				select {
				case entry, ok := <-dest:
					if !ok {
						fmt.Printf("Done %d after processing %d entries\n", id, count)
						return
					}
					fmt.Fprintln(f, entry)
					count++
				}
			}
		}(dest, i)
	}

	wg.Wait()
}
