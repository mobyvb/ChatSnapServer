package main

import (
	"context"
	"fmt"
	"log"

	model "github.com/mobyvb/ChatSnapServer/model/gen"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := "0.0.0.0:3000"

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := model.NewHelloServiceClient(conn)
	res, err := client.SayHello(context.Background(), &model.HelloRequest{})

	fmt.Println(res.GetResponse())
}
