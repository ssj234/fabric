package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"path"
)
import pb "github.com/hyperledger/awesomeProject/hello/hellopb"

var (
	serverkey = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/twoway" ,"peer0.server.key")
	servercrt  = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/twoway" ,"peer0.server.crt")
	cacrt  = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/twoway" ,"tlsca.org1.cmbc.com-cert.pem")
)
type HelloService struct {

}
func (s *HelloService) Hello(ctx context.Context,request *pb.Request) (*pb.Response, error){
	var req = request.Content
	var ret = "Hello " + req
	return &pb.Response{Content:ret},nil
}


func main()  {
	// 1.从文件读取私钥和证书创建TLS服务器
	certificate, err := tls.LoadX509KeyPair(servercrt, serverkey)
	if err != nil {
		log.Fatalf("could not load server key pair: %s", err)
	}
	// 2.根据cacrt创建就一个certificate pool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(cacrt)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}
	// 3.将cacrt加入到certificate pool
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append client certs")
	}

	// 4.创建credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

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
