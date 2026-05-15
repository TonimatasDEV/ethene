package network

import (
	"log"
	"net"
)

func InitReceiver() {
	listener, err := net.Listen("tcp", ":25565")

	if err != nil {
		log.Fatalf("Error listening: %v", err.Error())
	}

	defer listener.Close()

	log.Println("Listening on :25565")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting:", err)
			continue
		}

		session := NewConnection(conn)
		go func() {
			err := HandleConnection(session)
			if err != nil {
				log.Printf("Error handling connection: %v\n", err.Error())
			}
		}()
	}
}
