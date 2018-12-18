package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"path"
)
import pb "github.com/hyperledger/awesomeProject/hello/hellopb"

var (
	key = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/serverside" ,"private.key")
	servercrt  = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/serverside" ,"root.crt");
)
type HelloService struct {

}
func (s *HelloService) Hello(ctx context.Context,request *pb.Request) (*pb.Response, error){
	var req = request.Content
	var ret = "Hello " + req
	return &pb.Response{Content:ret},nil
}


func main()  {
	// 从文件读取私钥和证书创建TLS服务器
	creds, err := credentials.NewServerTLSFromFile(servercrt,key)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	var port int = 9999
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption // 服务器端的配置参数
	opts = append(opts,grpc.Creds(creds)) // 获取设置加密的ServerOption
	grpcServer := grpc.NewServer(opts...)
	// 创建一个Server，为其添加savedFeatures
	s := &HelloService{}
	pb.RegisterHelloServiceServer(grpcServer, s) // 注册到grpc服务器上
	grpcServer.Serve(lis) // 启动服务器

}
