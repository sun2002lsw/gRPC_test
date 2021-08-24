package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "helloWorldClient/pkg/helloWorld_protocolBuffer"
)

const address = "localhost:9999"

func main() {
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewSayHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, _ := client.GetHelloWorldFromServer(ctx, &pb.Empty{})
	fmt.Println("message from server:", response.Msg)

	time.Sleep(time.Minute)
}
