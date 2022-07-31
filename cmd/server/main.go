package main

import (
	"context"
	internal "github.com/antiartificial/book-service/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("Starting listening on port 8080")
	port := ":8080"

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Mongo Repository
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("testing")

	// create new mongo repository
	var repository internal.BookRepository = internal.NewMongoBookRepository(db)
	srv := internal.NewRPCServer(repository)

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
