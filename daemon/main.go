package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/zsxawerdu/pgrok/service"
	"google.golang.org/grpc"
)

const (
	PORT string = "18411"
)

type PipeService struct{}

func (p *PipeService) Write(ctx context.Context, w *service.WriteReq) (*service.WriteResp, error) {
	log.Printf("Pipe [%s] => Length Of Data %d\n", w.PipeName, len(w.Data))
	return &service.WriteResp{}, nil
}

func (p *PipeService) Read(ctx context.Context, r *service.ReadReq) (*service.ReadResp, error) {
	log.Printf("Pipe [%s] requested a read.", r.PipeName)
	resp := service.ReadResp{
		Data: []byte("hello from service"),
	}
	return &resp, nil
}

var pipe PipeService

func main() {
	l, err := net.Listen("tcp", "localhost:18411")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	service.RegisterPipeServer(server, &pipe)

	fmt.Println("Started Pipe Service")
	server.Serve(l)
}
