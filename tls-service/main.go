package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"path/filepath"
	"strings"

	pb "github.com/blablatov/stream-grpc/stream-tls-grpc/tls-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	crtFile            = filepath.Join("..", "certs", "server.crt")
	keyFile            = filepath.Join("..", "certs", "server.key")
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

const (
	port = ":50051"
)

func main() {
	// Считываем и анализируем открытый/закрытый ключи и создаем сертификат, чтобы включить TLS
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	opts := []grpc.ServerOption{
		// Enable TLS for all incoming connections.
		// Включаем TLS для всех входящих соединений, используя сертификаты для аутентификации
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		// Добавляем с помощью вызова grpc.UnaryInterceptor перехватчик, будет направлять все клиентские запросы
		grpc.UnaryInterceptor(ensureValidToken),
	}

	// Создаем новый экземпляр gRPC-сервера, передавая ему аутентификационные данные
	s := grpc.NewServer(opts...)
	// Регистрируем реализованный сервис на только что созданном gRPCсервере с помощью сгенерированных AP
	pb.RegisterProductInfoServer(s, &server{})

	lis, err := net.Listen("tcp", port) // Начинаем прослушивать TCP-соединение на порту 50051.
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting gRPC listener on port " + port)

	// Привязываем gRPC-сервер к прослушивателю и ждем появления сообщений на порту 50051.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Performs validation of token matching an arbitrary string.
	// Выполняем проверку токена, соответствующего произвольной строке.
	return token == "blablatok-tokblabla-blablatok"
}

// Определяем функцию ensureValidToken для проверки подлинности токена.
// Если тот отсутствует или недействителен, тогда перехватчик блокирует выполнение и возвращает ошибку.
// Или вызывается следующий обработчик, которому передается контекст и интерфейс.
func ensureValidToken(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}
