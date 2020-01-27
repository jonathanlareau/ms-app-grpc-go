/**
 * jonathan.lareau@gmail.com
 *
 * Very simple server example
 *
 * Listening on port 3000
 *
 **/

package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/jonathanlareau/ms-app-grpc-go/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var count = 0

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen on port 3000:  %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSayHelloServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(cxt context.Context, r *pb.SayHiRequest) (*pb.SayHelloResponse, error) {
	responseMsg := "Hello from go"

	msg := fmt.Sprintf("Msg: %s <-> %s -> %d", r.RequestMsg, responseMsg, count)

	result := &pb.SayHelloResponse{ResponseMsg: msg}

	log.Println(msg)

	count++

	return result, nil
}
