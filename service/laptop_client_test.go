package service_test

import (
	"context"
	"io"
	"net"
	"testing"

	"github.com/itzmanish/laptopstore/utils"

	"github.com/itzmanish/laptopstore/pb"
	"github.com/itzmanish/laptopstore/sample"
	"github.com/itzmanish/laptopstore/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t, service.NewInMemoryLaptopStore())
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedId := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Id, expectedId)

	// check laptop is saved on store or not
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check that saved laptop is same as we sent as request
	value, err := utils.CompareLaptop(laptop, other)
	require.NoError(t, err)
	require.True(t, value)
}

func TestClientFilterLaptop(t *testing.T) {
	t.Parallel()

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCore:  8,
		MinGhz:      2.2,
		MinRam: &pb.Memory{
			Unit:  pb.Memory_GIGABYTE,
			Value: 8,
		},
	}

	store := service.NewInMemoryLaptopStore()

	expectedId := make(map[string]bool)

	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()

		switch i {
		case 0:
			laptop.PriceUsd = 4000
		case 1:
			laptop.Cpu.Core = 6
		case 2:
			laptop.Cpu.MinGhz = 2.2
		case 3:
			laptop.Ram = &pb.Memory{
				Unit:  pb.Memory_GIGABYTE,
				Value: 4,
			}
		case 4:
			laptop.PriceUsd = 2000
			laptop.Cpu.Core = 12
			laptop.Cpu.MinGhz = 3
			laptop.Ram = &pb.Memory{
				Unit:  pb.Memory_GIGABYTE,
				Value: 16,
			}
			expectedId[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 2200
			laptop.Cpu.Core = 16
			laptop.Cpu.MinGhz = 3.5
			laptop.Ram = &pb.Memory{
				Unit:  pb.Memory_GIGABYTE,
				Value: 18,
			}
			expectedId[laptop.Id] = true
		}

		err := store.Save(laptop)
		require.NoError(t, err)
	}

	_, serverAddr := startTestLaptopServer(t, store)
	client := newTestLaptopClient(t, serverAddr)

	req := &pb.FilterLaptopRequest{
		Filter: filter,
	}
	stream, err := client.FilterLaptop(context.Background(), req)
	require.NoError(t, err)

	found := 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		require.Contains(t, expectedId, res.GetLaptop().GetId())

		found += 1
	}
}

func startTestLaptopServer(t *testing.T, store *service.InMemoryLaptopStore) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(store)

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}
