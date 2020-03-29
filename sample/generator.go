package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/itzmanish/laptopstore/pb"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new sample cpu
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	core := randomInt(2, 8)
	thread := randomInt(core, 16)
	minGhz := randomFloat64(2.0, 3.0)
	maxGhz := randomFloat64(minGhz, 4.0)
	cpu := &pb.CPU{
		Brand:  brand,
		Name:   name,
		Core:   uint32(core),
		Thread: uint32(thread),
		MinGhz: minGhz,
		MaxGhz: maxGhz,
	}
	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

// NewMemory return a new memory
func NewMemory() *pb.Memory {
	return &pb.Memory{
		Unit:  pb.Memory_GIGABYTE,
		Value: uint64(randomInt(2, 64)),
	}
}

// NewSSD returns a new ssd
func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Unit:  pb.Memory_GIGABYTE,
			Value: uint64(randomInt(128, 1024)),
		},
	}
}

// NewHDD returns a new HDD
func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Unit:  pb.Memory_TERABYTE,
			Value: uint64(randomInt(1, 6)),
		},
	}
}

// NewScreen returns a new Screen
func NewScreen() *pb.Screen {
	return &pb.Screen{
		SizeInch:   randomFloat32(24, 128),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
}

// NewLaptop returns new laptop
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Gpus:     []*pb.GPU{NewGPU()},
		Ram:      NewMemory(),
		Storages: []*pb.Storage{NewHDD(), NewSSD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2014, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
