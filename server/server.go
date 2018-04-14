package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-courses/freelance/api"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// StartGRPCServer для запуска grpc сервера
func StartGRPCServer(address string, s *api.Server) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance
	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the DoUser service to the server
	api.RegisterDoUsersServer(grpcServer, s)

	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}

// StartRESTServer для запуска rest сервера
func StartRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register DoUser endpoints
	err := api.RegisterDoUsersHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service InfoStatus: %s", err)
	}

	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)

	return nil
}
