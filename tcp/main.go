package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func pcrocessWorker(connections <-chan net.Conn, id int) {
	log.Printf("worker %d => up and ready", id)
	for conn := range connections {
		log.Printf("worker %d => Someone just connected to the magical server!", id)
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					log.Printf("worker %d => Someone just disconnected", id)
				}
				break
			}

			log.Printf("worker %d => received: %q", id, buf[:n])
		}
		conn.Close()
	}
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server is up and running on %s", lis.Addr())

	connectionsChan := make(chan net.Conn, 10)
	for i := 1; i <= 10; i++ {
		go pcrocessWorker(connectionsChan, i)
	}

	for {
		connection, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		connectionsChan <- connection
		fmt.Println("test")
	}
}
