package main

//use openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 -nodes 
// to generate key.pem and cert.pem

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"
	"github.com/quic-go/quic-go"
)

func main() {
	tlsConfig := generateTLSConfig()

	// Start a QUIC listener on port 4242
	listener, err := quic.ListenAddr(":4242", tlsConfig, nil)
	if err != nil {
		log.Fatalf("Failed to start QUIC listener: %v", err)
	}
	fmt.Println("Server is listening on port 4242")

	for {
		// Pass a context to the Accept method
		ctx := context.Background()
		conn, err := listener.Accept(ctx)
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn quic.Connection) {
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Printf("Failed to accept stream: %v", err)
		return
	}

	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		log.Printf("Failed to read from stream: %v", err)
		return
	}

	fmt.Printf("Received: %s\n", string(buf[:n]))
	n, err=stream.Write([]byte("Hello, client!"))
	fmt.Printf("send: %d\n", n)
	time.Sleep(2 * time.Second)
	defer stream.Close()
}

func generateTLSConfig() *tls.Config {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Failed to load certificates: %v", err)
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

