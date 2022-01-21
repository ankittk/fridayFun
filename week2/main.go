package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	pb "weavelab.xyz/week2/proto/frequency"
)

type server struct {
	pb.UnimplementedWordFrequencyServer
}

func (s *server) Calculate(_ context.Context, req *pb.InputRequest) (*pb.OutputResponse, error) {
	input := req.Text
	splitWord := strings.Fields(input)
	wordCount := make(map[string]int64)
	for _, value := range splitWord {
		_, alreadyExist := wordCount[value]
		if alreadyExist {
			wordCount[value] += 1
		} else {
			wordCount[value] = 1
		}
	}
	for i, j := range wordCount {
		fmt.Printf("Word is : %v and its count is %v: ", i, j)
		fmt.Println()
	}
	return &pb.OutputResponse{
		Result: wordCount,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterWordFrequencyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("Hosting server on: %s", lis.Addr().String())
}
