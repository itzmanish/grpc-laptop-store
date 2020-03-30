package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/itzmanish/laptopstore/pb"
	"github.com/itzmanish/laptopstore/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	address := flag.String("address", "", "Address of server")
	flag.Parse()
	log.Printf("Connecting to server: %v", *address)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("Laptop already exist")
		} else {
			log.Fatalf("Unable to create laptop: %v", err)
		}
		return
	}
	log.Printf("Created laptop with id: %v", res.Id)
}
