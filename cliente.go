package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"os"

	"github.com/quic-go/quic-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: quic_client <server_ip>")
		return
	}

	serverIP := os.Args[1]

	// Cria um contexto de fundo
	ctx := context.Background()

	// Estabelece uma conexão QUIC com o servidor
	session, err := quic.DialAddr(ctx, fmt.Sprintf("%s:4242", serverIP), &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer session.CloseWithError(0, "Closing connection")

	// Abre um stream sincronizado
	stream, err := session.OpenStreamSync(ctx)
	if err != nil {
		fmt.Println("Stream error:", err)
		return
	}
	defer stream.Close()

	// Cria o arquivo para salvar o conteúdo recebido
	file, err := os.Create("received.html")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// Copia os dados recebidos no stream para o arquivo
	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println("Failed to receive file:", err)
		return
	}

	fmt.Println("File received successfully!")
}
