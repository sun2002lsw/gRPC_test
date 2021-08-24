package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "helloWorldServer/pkg/helloWorld_protocolBuffer"
)

const port = ":9999"

func main() {
	listen, _ := net.Listen("tcp", port)

	s := grpc.NewServer()
	pb.RegisterSayHelloServer(s, &server{})

	fmt.Println("Start serving...")
	s.Serve(listen)
}
