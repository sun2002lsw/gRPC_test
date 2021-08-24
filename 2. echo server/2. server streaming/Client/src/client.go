package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "echo/pkg/echo"
)

const address = "localhost:9999"

func main() {
	consoleInput := bufio.NewReader(os.Stdin)

	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	client := pb.NewEchoServiceClient(conn)
	defer conn.Close()

	for input := getInput(consoleInput); input != "exit"; input = getInput(consoleInput) {
		printEcho(client, input)
	}
}

func getInput(consoleInput *bufio.Reader) string {
	fmt.Print("input: ")
	text, _ := consoleInput.ReadString('\n')

	for lastChar := text[len(text) - 1]; lastChar == '\n' || lastChar == '\r'; lastChar = text[len(text) - 1] {
		text = text[:len(text) - 1]
	}

	return text
}

func printEcho(client pb.EchoServiceClient, input string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, _ := client.Echo(ctx, &pb.Word{W: input})
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		fmt.Println("message from server:", res.W)
	}

	fmt.Println()
}
