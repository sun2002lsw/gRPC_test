package main

import (
	"context"
	"fmt"

	pb "echoServer/pkg/echo"
)

type server struct {}

func (s *server) Echo(ctx context.Context, in *pb.Word) (out *pb.Word, err error) {
	fmt.Println("message from client:", in.W)

	out = new (pb.Word)
	out.W = in.W

	return out, nil
}
