package main

import (
	"log"
	"net"

	alert "github.com/Rconduru/React-native_gRPC-web/server/api"
	alertPb "github.com/Rconduru/React-native_gRPC-web/server/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := alert.Server{}

	alertPb.RegisterAlertServiceServer(grpcServer, &server)

	reflection.Register(grpcServer)
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
