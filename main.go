package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 8)
	for err == nil {
		_, err = file.Read(data)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Fatal(err)
		}
		fmt.Printf("read: %s\n", data)
	}
}
