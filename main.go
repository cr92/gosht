package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/cr92/gosht/customer"
	"github.com/cr92/gosht/dataSrc"
	"github.com/cr92/gosht/file"
)

func main() {
	loadEnv()
	dataFile := file.File{Path: os.Getenv("DATA")}

	start := time.Now()
	processData(&dataFile)
	fmt.Println(time.Since(start))
}

func processData(ds dataSrc.DataSrc) {
	ctx, cancel := context.WithTimeout(context.Background(), 450*time.Millisecond)
	defer cancel()
	var wg sync.WaitGroup

	dest := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ds.ReadLine(ctx, dest)
		close(dest)
	}()

	for i := range 10 {
		wg.Add(1)
		go processLine(ctx, dest, i, &wg)
	}
	wg.Wait()
}

func processLine(ctx context.Context, dest chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
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

	for {
		select {
		case entry, ok := <-dest:
			if !ok {
				fmt.Printf("Done %d after processing %d entries\n", id, count)
				return
			}
			j, _ := json.MarshalIndent(customer.CreateCustomer(entry), "", "    ")
			fmt.Fprintln(f, string(j))
			count++
		case <-ctx.Done():
			fmt.Printf("Consumer %d cancelled\n", id)
			return
		}
	}
}
