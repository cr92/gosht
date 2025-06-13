package file

import (
	"bufio"
	"log"
	"os"
)

type File struct {
	Path string
}

func (f *File) ReadLine(dest chan string) {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Fatal("Opening file failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dest <- scanner.Text()
	}
	return
}
