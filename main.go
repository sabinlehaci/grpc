//server app
package main

import (
	"log"
	"net"
	usersv1 "github.com/sabinlehaci/grpc-test/gen/go/users/v1"
	users "github.com/sabinlehaci/grpc-test/users"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}

func run() error {
	srv := grpc.NewServer()

	usersv1.RegisterUserServiceServer(srv, users.NewUserService())

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	log.Println("Serving on :8080")
	return srv.Serve(listener)
}
