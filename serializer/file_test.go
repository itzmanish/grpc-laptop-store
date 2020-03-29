package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/itzmanish/laptopstore/pb"
	"github.com/itzmanish/laptopstore/sample"
	"github.com/itzmanish/laptopstore/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobuffToBinary(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobuffFromBinary(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobuffToJSON(laptop1, jsonFile)
	require.NoError(t, err)
}
