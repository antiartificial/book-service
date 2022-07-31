package internal

import (
	api "github.com/antiartificial/book-service/api/v1"
	"google.golang.org/grpc"
)

type grpcServer struct {
	BookRepository BookRepository
	api.UnimplementedBookServiceServer
}

func NewRPCServer(repository BookRepository) *grpc.Server {
	srv := grpcServer{
		BookRepository: repository,
	}
	gsrv := grpc.NewServer()
	api.RegisterBookServiceServer(gsrv, &srv)
	return gsrv
}
