package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func run_rename() {
	filename := "output.txt"
	data := []byte("Hello db")
	n, err := rename(filename, data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Wrote %d bytes to file %s\n", n, filename)
}

func rename(filename string, data []byte) (int, error) {
	randInt := rand.Int()
	tmp := fmt.Sprintf("%s.tmp.%d", filename, randInt)

	file, err := os.OpenFile(tmp, os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0664)
	defer func() {
		file.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	n, err := file.Write(data)
	if err != nil {
		return 0, err
	}

	err = file.Sync()
	if err != nil {
		return 0, err
	}

	return n, os.Rename(tmp, filename)
}
