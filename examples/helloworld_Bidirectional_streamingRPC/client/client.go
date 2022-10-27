package main

import (
	pb "HelloWorld_BRPC/proto"
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
	// name := &[]string{}
	name := "sss"
	_ = SayRecord(client, &pb.HelloRequest{Name: name})
}

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 7; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("client")
		fmt.Println(resp)

	}
	_ = stream.CloseSend()
	return nil
}
