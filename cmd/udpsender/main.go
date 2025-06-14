package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:42069")
	if err != nil {
		log.Fatalf("Error resolving UDP adress: %v", err)
	}

	connection, err := net.DialUDP("udp", nil, udpAddress)
	if err != nil {
		log.Fatalf("error creating UDP connection: %v\n", err)
	}
	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Printf("Error reading line from standard input: %v\n", err)
		}

		_, err = connection.Write(([]byte)(line))
		if err != nil {
			fmt.Printf("Error writing line though UDP connection: %v\n", err)
		}

	}
}
