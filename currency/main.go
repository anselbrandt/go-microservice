package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/anselbrandt/go-microservice/currency/data"

	protos "github.com/anselbrandt/go-microservice/currency/protos"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 9092, "The server port")
	log  = hclog.Default()
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	rates *data.ExchangeRates
	log   hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(p protos.UnimplementedCurrencyServer, r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{p, r, l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	gs := grpc.NewServer()

	protos.RegisterCurrencyServer(gs, &Currency{})

	reflection.Register(gs)

	gs.Serve(lis)
}
