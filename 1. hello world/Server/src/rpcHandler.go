package main

import (
	"context"
	"fmt"

	pb "helloWorldServer/pkg/helloWorld_protocolBuffer"
)

type server struct {}

func (s *server) GetHelloWorldFromServer(ctx context.Context, in *pb.Empty) (*pb.HelloMsg, error) {
	fmt.Println("function called")

	ret := pb.HelloMsg{}
	ret.Msg = "Hello World"

	return &ret, nil
}
