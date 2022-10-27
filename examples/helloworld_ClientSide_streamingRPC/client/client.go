package main

import (
	pb "HelloWorld_CRPC/proto"
	"context"
	"flag"

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
	}
	resp, _ := stream.CloseAndRecv()
	println(resp)
	return nil
}
