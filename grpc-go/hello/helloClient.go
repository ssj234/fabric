package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "github.com/hyperledger/awesomeProject/hello/hellopb"
)

func main()  {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:9999", opts...)
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