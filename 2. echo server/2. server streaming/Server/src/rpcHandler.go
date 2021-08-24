package main

import (
	"fmt"
	"strconv"

	pb "echoServer/pkg/echo"
)

type server struct {}

func (s *server) Echo(in *pb.Word, stream pb.EchoService_EchoServer) error {
	fmt.Println("message from client:", in.W)

	for idx, char := range in.W {
		res := "[" + strconv.Itoa(idx) + " - " + string(char) + "]"

		stream.Send(&pb.Word{W: res})
	}

	return nil
}
