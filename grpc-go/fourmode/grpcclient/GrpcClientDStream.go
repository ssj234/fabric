package main

import (
	"context"
	pb "github.com/hyperledger/awesomeProject/fourmode/grpcguide"
	"google.golang.org/grpc"
	"io"
	"log"
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

	stream, err := client.RouteChat(context.Background())
	// 由于在发送RouteNode的同时，还要接收RouteNode，因此使用go
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc) // 关闭chan
				return
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()

	for _, note := range getNodes() {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err.Error())
		}
	}
	stream.CloseSend()
	<-waitc // 等待chan关闭
}

func getNodes() []*pb.RouteNote{
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	return notes
}