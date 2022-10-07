package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server is up and running on %s", lis.Addr())
	for {
		accept, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			defer c.Close()
			log.Println("Someone just connected to the magical server!")
			buf := make([]byte, 1024)
			for {
				n, err := c.Read(buf)
				if err != nil {
					if err == io.EOF {
						log.Println("Someone just disconnected")
					}
					return
				}

				log.Printf("received: %q", buf[:n])
			}
		}(accept)
	}
}
