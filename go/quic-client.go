package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"
	"github.com/quic-go/quic-go"
)

func main() {
	// Create a TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // WARNING: For development/testing only
	}

	// Create a context for the connection
	ctx := context.Background()

	// Connect to the QUIC server
	conn, err := quic.DialAddr(ctx, "localhost:4242", tlsConfig, nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.CloseWithError(0, "client shutting donw")

	// Open a stream
	stream, err := conn.OpenStreamSync(ctx)
	if err != nil {
		log.Fatalf("Failed to open stream: %v", err)
	}

	// Send a message to the server
	message := "Hello, server!"
	stream.Write([]byte(message))
	fmt.Printf("Sent: %s\n", message)

	// Receive a response from the server
	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		log.Fatalf("Failed to read from stream: %v", err)
	}
	fmt.Printf("Received: %s\n", string(buf[:n]))

	defer stream.Close()
}

