package file

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
)

type File struct {
	Path string
}

func (f *File) ReadLine(ctx context.Context, dest chan string) {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Fatal("Opening file failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case dest <- scanner.Text():
		case <-ctx.Done():
			fmt.Println("Producer cancelled")
			return
		}
	}
	return
}
