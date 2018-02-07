package main

import (
	"fmt"
	"log"
	"net"

	proto "github.com/muhfaris/goFun/basic_serialdata/03_go_gRPC/server/kittens"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type kittenServer struct{}

func (k *kittenServer) Hello(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	response := &proto.Response{}
	response.Name = fmt.Sprintf("Hello, my name is %v:", request.Name)

	return response, nil
}
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterKittensServer(grpcServer, &kittenServer{})
	grpcServer.Serve(lis)

}
