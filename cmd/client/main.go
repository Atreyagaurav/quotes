package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	api "github.com/atreyagaurav/quotes/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	address := flag.String("server", "", "gRPC server in format host:port")
	genre := flag.String("genre", "", "genre of the quote to return")
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := api.NewQuotesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := api.QuoteRequest{
		Id:    1,
		Genre: *genre,
	}
	res, err := c.Get(ctx, &req)
	fmt.Printf(res.Body)
}
