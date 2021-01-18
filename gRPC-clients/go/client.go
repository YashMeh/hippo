package main

/**
This is a sample gRPC client to talk to the hippo service
**/

import (
	"context"
	"log"

	pb "github.com/mailgun/kafka-pixy/gen/golang"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("127.0.0.1:19091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()
	//Thread safe client
	client := pb.NewKafkaPixyClient(conn)
	messageLog := []byte("I ate a clock yesterday, it was very time-consuming.")
	//Request structure https://github.com/mailgun/kafka-pixy
	rs, err := client.Produce(context.Background(), &pb.ProdRq{
		Topic: "foo", Message: messageLog})
	if err != nil {
		panic(err)
	}
	log.Printf("Produced: partition=%d, offset=%d\n", rs.Partition, rs.Offset)
}
