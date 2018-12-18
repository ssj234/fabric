package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
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

	// 调用ListFeatures
	point1 := &pb.Point{Longitude:-743977337,Latitude:407033786}
	point2 := &pb.Point{Longitude:-740477477,Latitude:414653148}
	rectangle := &pb.Rectangle{Lo:point1,Hi:point2}
	stream, err := client.ListFeatures(context.Background(),rectangle)
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}
