package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"github.com/itzmanish/laptopstore/pb"
	"github.com/itzmanish/laptopstore/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createLaptopRequest(laptopClient pb.LaptopServiceClient) {

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

func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	req := &pb.FilterLaptopRequest{
		Filter: filter,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := laptopClient.FilterLaptop(ctx, req)
	if err != nil {
		log.Fatalf("Unable to search laptop: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Unable to recieve response from server: %v", err)
		}
		laptop := res.GetLaptop()
		log.Printf("-found: %v", laptop.GetId())
		log.Printf(" + brand: %v", laptop.GetBrand())
		log.Printf(" + name: %v", laptop.GetName())
		log.Printf(" + cpu core: %v", laptop.GetCpu().GetCore())
		log.Printf(" + cpu min ghz: %v", laptop.GetCpu().GetMinGhz())
		log.Printf(" + ram: %v", laptop.GetRam())
		log.Printf(" + price: %v", laptop.GetPriceUsd())
	}

}

func main() {
	address := flag.String("address", "", "Address of server")
	flag.Parse()
	log.Printf("Connecting to server: %v", *address)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	for i := 0; i < 10; i++ {
		createLaptopRequest(laptopClient)
	}

	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCore:  4,
		MinGhz:      2.0,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	searchLaptop(laptopClient, filter)
}
