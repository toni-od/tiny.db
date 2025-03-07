package main

import (
	"fmt"
	"log"
	"os"
)

func Write(filename string, data string) (int, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n, err := file.Write([]byte(data))
	if err != nil {
		return 0, err
	}

	return n, file.Sync()
}

func run_writer() {
	filename := "output.txt"
	data := "Hello db!"
	n, err := Write(filename, data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Wrote %d bytes to file %s\n", n, filename)
}
