package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)
import pb "github.com/hyperledger/awesomeProject/hello/hellopb"


type HelloService struct {

}
func (s *HelloService) Hello(ctx context.Context,request *pb.Request) (*pb.Response, error){
	var req = request.Content
	var ret = "Hello " + req
	return &pb.Response{Content:ret},nil
}


func main()  {
	var port int = 9999
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption // 服务器端的配置参数
	grpcServer := grpc.NewServer(opts...)
	// 创建一个Server，为其添加savedFeatures
	s := &HelloService{}
	pb.RegisterHelloServiceServer(grpcServer, s) // 注册到grpc服务器上
	grpcServer.Serve(lis) // 启动服务器

}
