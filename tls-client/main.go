package main

import (
	"context"
	"log"
	"path/filepath"
	"time"

	pb "github.com/blablatov/stream-tls-grpc/tls-proto"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

var (
	crtFile = filepath.Join("..", "certs", "server.crt")
)

const (
	address = "localhost:50051"
	//address  = "net-tls-service:50051"
	hostname = "localhost"
)

func main() {
	// Set up the credentials for the connection.
	// Значение токена OAuth2. Используем строку, прописанную в коде.
	tokau := oauth.NewOauthAccess(fetchToken())

	// Считываем и анализируем публичный сертификат, чтобы включить TLS.
	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}

	// Указываем аутентификационные данные для транспортного протокола с помощью DialOption.
	opts := []grpc.DialOption{
		// Указываем один и тот же токен OAuth в параметрах всех вызовов в рамках одного соединения.
		// Если нужно указывать токен для каждого вызова отдельно, используем CallOption.
		grpc.WithPerRPCCredentials(tokau),
		// transport credentials.
		grpc.WithTransportCredentials(creds),
	}

	// Set up a connection to the server.
	// Устанавливаем безопасное соединение с сервером, передавая параметры аутентификации
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Передаем соединение и создаем заглушку.
	// Ее экземпляр содержит все удаленные методы, которые можно вызвать на сервере.
	c := pb.NewProductInfoClient(conn)

	// Contact the server and print out its response.
	name := "Sumsung S99"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2023"
	price := float32(7000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Println("Product: ", product.String())
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "blablatok-tokblabla-blablatok",
	}
}
