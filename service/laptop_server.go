package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/itzmanish/laptopstore/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer struct consist of Store method
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopServer returns new Laptop proto message
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

// CreateLaptop returns a CreateLaptopResponse and error
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("Recieving laptop request with id : %v", laptop.GetId())
	if len(laptop.GetId()) > 0 {
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid : %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Unable to generate id for laptop: %v", err)
		}
		laptop.Id = id.String()
	}

	// halt for some second
	// time.Sleep(6 * time.Second)

	if ctx.Err() == context.Canceled {
		log.Printf("Request canceled from client")
		return nil, status.Errorf(codes.Canceled, "Request is canceled")
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline exceeded.")
		return nil, status.Errorf(codes.DeadlineExceeded, "Deadline is exceeded")
	}

	// save the laptop to store
	if err := server.Store.Save(laptop); err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExist) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "Unable to save to store: %v", err)
	}

	return &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}, nil
}

// FilterLaptop fiilters laptop and send a stream and return and error
func (server *LaptopServer) FilterLaptop(req *pb.FilterLaptopRequest, stream pb.LaptopService_FilterLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("recieved a search laptop request with filter: %v", filter)

	err := server.Store.Search(
		stream.Context(),
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.FilterLaptopResponse{
				Laptop: laptop,
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("laptop with id %v sent", laptop.Id)
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error occured during filter: %v", err)
	}
	return nil
}
