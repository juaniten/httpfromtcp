package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		log.Fatalf("Error opening TCP listener: %v", err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting TCP connection: %v", err)
		}
		fmt.Println()
		fmt.Println("TCP connection accepted")
		fmt.Println("-----------------------")
		linesChannel := getLinesChannel(connection)
		for line := range linesChannel {
			fmt.Println(line)
		}
		fmt.Println("-----------------------")
		fmt.Println("TCP connection closed")
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
