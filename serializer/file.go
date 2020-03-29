package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// WriteProtobuffToBinary serialize protobuff to binary or returns an error
func WriteProtobuffToBinary(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("error on marshalling %v", err)
	}
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("errorn on writing to file: %v", err)
	}
	return nil
}

// ReadProtobuffFromBinary deserialize binary to protobuff or returns an error
func ReadProtobuffFromBinary(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Unable to read from binary: %v", err)
	}
	if err := proto.Unmarshal(data, message); err != nil {
		return fmt.Errorf("unable to convert Binary to protobuffer: %v", err)
	}
	return nil
}

//WriteProtobuffToJSON writes protobuff to json file and returns error
func WriteProtobuffToJSON(message proto.Message, filename string) error {
	data, err := ProtobuffToJSON(message)
	if err != nil {
		return fmt.Errorf("Unable to convert protobuff to json string: %v", err)
	}

	if err := ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
		return fmt.Errorf("Unable to write json string to file: %v", err)
	}
	return nil
}
