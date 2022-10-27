package main

import (
	pb "HelloWorld_CRPC/proto"
	"flag"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

var port string

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			message := &pb.HelloReply{Message: "sss"}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}
		fmt.Println(resp)
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
