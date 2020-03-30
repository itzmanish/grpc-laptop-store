package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/itzmanish/laptopstore/pb"
	"google.golang.org/grpc"

	"github.com/itzmanish/laptopstore/service"
)

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Printf("server started on port: %v", *port)

	server := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Unable to listen tcp connection: %v", err)
	}

	pb.RegisterLaptopServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}

}
