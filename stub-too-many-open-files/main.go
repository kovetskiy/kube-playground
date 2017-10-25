package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("opening too many open files")

	var files []*os.File
	var i int64
	for {
		file, err := os.OpenFile(fmt.Sprint(i), os.O_CREATE, 0644)
		if err != nil {
			log.Println(err)
			continue
		}

		files = append(files, file)

		log.Printf("opened %d", len(files))
		i++
	}
}
