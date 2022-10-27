package main

import (
	pb "HelloWorld_SRPC/proto"
	"context"
	"flag"
	"fmt"
	"io"

	"google.golang.org/grpc"
)

var port string

func main() {
	flag.Parse()
	port = "8000"
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	// , &pb.HelloRequest{Name: *name}
	name := &[]string{}
	_ = SayHello_1(client, &pb.HelloRequest{Name: *name})
}

func SayHello_1(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println("resp: %v", resp)
	}
	return nil
}
