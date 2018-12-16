package main

import (
    "context"
    "fmt"
    "github.com/alfredd/myfirstgoproject/protocol"
    "google.golang.org/grpc"
    "log"
    "net"
)

type Server struct{}

func (s *Server) Echo(ctx context.Context, e *echo.EchoMessage) (*echo.EchoMessage, error) {
    return e, nil
}

func main() {
    port := 4550
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
    if err != nil {
        log.Fatal("Error in listening to port")
    }

    server := grpc.NewServer()
    echo.RegisterEchoServerServer(server, &Server{})
    server.Serve(lis)
}
