package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/atreyagaurav/quotes/pkg/protocol/grpc"
	"github.com/atreyagaurav/quotes/pkg/service"
)

type Config struct {
	GRPCport string
}

func RunServer() error {
	ctx := context.Background()
	var cfg Config
	flag.StringVar(&cfg.GRPCport, "port", "string", "Port for gRCP to bind")
	flag.Parse()
	if len(cfg.GRPCport) == 0 {
		return fmt.Errorf("Invalid port given:%s", cfg)
	}
	name := "at"
	server := service.NewQuotesServer(&name)
	return grpc.RunServer(ctx, server, cfg.GRPCport)
}
