package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/quic-go/quic-go"
)

func main() {
	listener, err := quic.ListenAddr("0.0.0.0:4242", generateTLSConfig(), nil)
	if err != nil {
		fmt.Println("Failed to start QUIC server:", err)
		return
	}
	fmt.Println("QUIC server listening on port 4242")

	for {
		session, err := listener.Accept(nil)
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleSession(session)
	}
}

func handleSession(session quic.Connection) {
	stream, err := session.AcceptStream(nil)
	if err != nil {
		fmt.Println("Stream error:", err)
		return
	}
	defer stream.Close()

	file, err := os.Open("index.html") // Serve o arquivo index.html
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(stream, file)
	if err != nil {
		fmt.Println("Failed to send file:", err)
	}
}

func generateTLSConfig() *tls.Config {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"quic-example"},
	}
}
