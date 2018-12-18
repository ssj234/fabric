package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "github.com/hyperledger/awesomeProject/hello/hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"path"
)


var (
	clientkey = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/twoway" ,"peer1.server.key")
	clientcrt  = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/twoway" ,"peer1.server.crt")
	clientcacrt  = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/twoway" ,"tlsca.org1.cmbc.com-cert.pem")
)
func main()  {
	// 1.从磁盘加载证书
	certificate, err := tls.LoadX509KeyPair(clientcrt, clientkey)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}
	// 2.创建证书pool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(clientcacrt)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// 3.将cacrt加入证书pool
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certs")
	}
	// 4.配置credentials
	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "peer0.org1.cmbc.com", // NOTE: this is required!
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial("peer0.org1.cmbc.com:9999", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)

	// 调用GetFeature
	response, err := client.Hello(context.Background(),&pb.Request{Content:"shisj"})
	if err != nil {
		log.Fatalf("%v.Hello(_) = _, %v: ", client, err)
	}
	println(response.Content)
}