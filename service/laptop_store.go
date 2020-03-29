package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/itzmanish/laptopstore/pb"
	"github.com/jinzhu/copier"
)

// ErrAlreadyExist error
var ErrAlreadyExist = errors.New("Laptop already exist")

// LaptopStore interface consist Save method
type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

// InMemoryLaptopStore struct have map of string and laptop as data
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInMemoryLaptopStore return new InMemoryLaptopStore object
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save method saves laptop to InMemoryLaptopStore and returns an error
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExist
	}

	// deep copy
	other := &pb.Laptop{}
	if err := copier.Copy(other, laptop); err != nil {
		return fmt.Errorf("Unable to copy laptop: %v", err)
	}

	store.data[other.Id] = other
	return nil
}
