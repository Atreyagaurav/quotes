package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	api "github.com/atreyagaurav/quotes/pkg/api"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, ap api.QuotesServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	api.RegisterQuotesServer(server, ap)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Shutting down")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	fmt.Println("Starting Server...")
	return server.Serve(listen)
}
