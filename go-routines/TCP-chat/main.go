package main

import (
	"log"
	"net"
)

func main() {
	s := newSever()
	go s.run()


	listener, err := net.Listen("tcp", ":22")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}


	defer listener.Close()
	log.Printf("Server is running on :22")



	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err.Error())
			continue
		}
		
		go s.newClient(conn)
	}
}
