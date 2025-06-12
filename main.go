package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filePath = "./messages.txt"

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ch := getLinesChannel(f)

	for line := range ch {
		fmt.Println("read:", line)
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {

	ch := make(chan string)
	go func() {

		defer f.Close()
		defer close(ch)
		currentLineContents := ""

		for {
			buffer := make([]byte, 8, 8)
			n, err := f.Read(buffer)

			if err != nil {
				if currentLineContents != "" {
					ch <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}

			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				ch <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}

			currentLineContents += parts[len(parts)-1]
		}

	}()
	return ch
}
