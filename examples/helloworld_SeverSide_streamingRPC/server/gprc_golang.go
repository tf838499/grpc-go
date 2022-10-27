package main

import (
	pb "HelloWorld_SRPC/proto"
	"flag"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

var port string

func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 1000; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello.list" + string(n)})
	}
	return nil
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
