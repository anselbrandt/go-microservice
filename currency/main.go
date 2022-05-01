package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	protos "github.com/anselbrandt/go-microservice/currency/protos"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 9092, "The server port")
	log  = hclog.Default()
)

type server struct {
	protos.UnimplementedCurrencyServer
}

func (s *server) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	gs := grpc.NewServer()

	protos.RegisterCurrencyServer(gs, &server{})

	reflection.Register(gs)

	gs.Serve(lis)
}
