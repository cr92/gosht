package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	start := time.Now()
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	nc := File{os.Getenv("DATA")}
	dest := make(chan string)
	done := make(chan bool)
	defer close(done)
	defer close(dest)

	go nc.Spew(dest, done)

	for {
		select {
		case m := <-dest:
			fmt.Println(m)
		case <-done:
			fmt.Println("Done")
			fmt.Println(time.Since(start))
			return
		}
	}
}

type File struct {
	Path string
}

func (c *File) Spew(dest chan string, done chan bool) {
	f, e := os.Open(c.Path)
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		dest <- sc.Text()
	}
	done <- true
}

type C interface {
	Spew() string
}
