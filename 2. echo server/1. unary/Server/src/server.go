package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "echoServer/pkg/echo"
)

const port = ":9999"

func main() {
	listen, _ := net.Listen("tcp", port)

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})

	fmt.Println("Start serving...")
	s.Serve(listen)
}
