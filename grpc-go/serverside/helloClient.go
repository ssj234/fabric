package main

import (
	"context"
	pb "github.com/hyperledger/awesomeProject/hello/hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"path"
)


func main()  {
	var servercrt = path.Join("/home/shisj/mycode/gowksp/src/github.com/hyperledger/awesomeProject/serverside" ,"root.crt");
	creds, err := credentials.NewClientTLSFromFile(servercrt, "")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial("jianshu.test.com:9999", opts...)
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