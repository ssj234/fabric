package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "github.com/hyperledger/awesomeProject/fourmode/grpcguide"
)

func main()  {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:9999", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn) // 使用pb生成client

	// 调用GetFeature
	feature, err := client.GetFeature(context.Background(), &pb.Point{Latitude:409146138, Longitude:-746188906})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}
