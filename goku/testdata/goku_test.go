package testdata

import (
	"bytes"
	"github.com/v1990/protoc-gen-goku/goku"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	_ "github.com/v1990/protoc-gen-goku/plugins"
)

func TestRun(t *testing.T) {
	req := mockRequestFromProtoJSONFile("descriptor.proto.json")
	parameter := strings.Join([]string{
		"debug=1",
		"conf=../../examples/goku/config.yaml",
		"workPath=../../examples/goku",
		"outPath=../../examples/goku/out",
		"outPkgName=descriptors",
	}, ",")
	req.Parameter = &parameter

	stdin := mockStdIn(req)
	goku.Run(stdin, os.Stdout, os.Stderr)
}

func mockStdIn(req *pluginpb.CodeGeneratorRequest) io.Reader {
	data, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(data)
}

func mockRequestFromProtoJSONFile(filename string) *pluginpb.CodeGeneratorRequest {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	req := &pluginpb.CodeGeneratorRequest{}

	err = protojson.Unmarshal(content, req)
	if err != nil {
		panic(err)
	}
	return req
}
