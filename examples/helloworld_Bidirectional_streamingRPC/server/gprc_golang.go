package main

import (
	pb "HelloWorld_BRPC/proto"
	"flag"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

var port string

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "say.route"})
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		fmt.Println("server")
		fmt.Println(resp)
	}
}

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊")
	flag.Parse()
}

func main() {

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)

}
