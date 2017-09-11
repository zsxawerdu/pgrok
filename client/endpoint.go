package main

import (
	"context"
	"log"

	"github.com/zsxawerdu/pgrok/service"
	"google.golang.org/grpc"
)

type RequestWriter interface {
	WriteRequest(service.WriteReq) (int, error)
}

var DefaultPipeClient service.PipeClient = createPipeClient()

type EndpointWriter struct {
	service.PipeClient
	ctx context.Context
}

func NewEndpointWriter(client service.PipeClient) *EndpointWriter {
	if client == nil {
		client = DefaultPipeClient
	}
	e := EndpointWriter{client, context.Background()}
	return &e
}

func (e *EndpointWriter) WriteRequest(req service.WriteReq) (int, error) {
	resp, err := e.PipeClient.Write(e.ctx, &req)
	log.Println(resp, err)
	return len(req.Data), nil
}

func createPipeClient() service.PipeClient {
	conn, err := grpc.Dial("localhost:18411", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return service.NewPipeClient(conn)

}
