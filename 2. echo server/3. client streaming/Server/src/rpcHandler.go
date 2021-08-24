package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	pb "echoServer/pkg/echo"
)

type server struct {}

func (s *server) Echo(stream pb.EchoService_EchoServer) error {
	var words []string

	for {
		word, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Word{W: getWordsInfo(words)})
		}

		fmt.Println("message from client:", word.W)
		words = append(words, word.W)
	}
}

func getWordsInfo(words []string) string {
	wordCnt := strconv.Itoa(len(words))
 	fullWord := strings.Join(words, " ")

	return "[wordCnt: " + wordCnt + ", fullWord: " + fullWord + "]"
}
