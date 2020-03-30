package utils

import (
	"fmt"
	"reflect"

	"github.com/itzmanish/laptopstore/pb"
	"github.com/itzmanish/laptopstore/serializer"
)

// CompareLaptop compare two laptop and return a boolean
func CompareLaptop(laptop, other *pb.Laptop) (bool, error) {
	laptop1, err := serializer.ProtobuffToJSON(laptop)
	if err != nil {
		return false, fmt.Errorf("Unable to convert proto message: %v", err)
	}
	laptop2, err := serializer.ProtobuffToJSON(other)
	if err != nil {
		return false, fmt.Errorf("Unable to convert proto message: %v", err)
	}
	return reflect.DeepEqual(laptop1, laptop2), nil
}
