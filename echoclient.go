package main

import (
	"context"
	"fmt"
	"github.com/alfredd/myfirstgoproject/protocol"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	port := 4550
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error occurred while connecting to server: %v", err)
	}
	defer conn.Close()
	client := echo.NewEchoServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := client.Echo(ctx, &echo.EchoMessage{Message: "Whats up!"})
	if err != nil {
		log.Fatalf("Error occurred while sending message to server: %v", err)
	}
	log.Printf("Response from server: %s", reply.Message)

}
