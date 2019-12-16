package service

import (
	// "json"

	"context"
	"fmt"
	"net/http"
	"time"

	api "github.com/atreyagaurav/quotes/pkg/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var errMsg = "Server not connected"

func All(c *gin.Context) {
	genre := "love"
	c.HTML(http.StatusOK, "tem.html", gin.H{
		"str": *GetResp(&genre),
	})
}

func Random(c *gin.Context) {
	genre := "random"
	c.HTML(http.StatusOK, "tem.html", gin.H{
		"str": *GetResp(&genre),
	})
}

func Genre(c *gin.Context) {
	genre := c.Param("reqgenre")
	c.HTML(http.StatusOK, "tem.html", gin.H{
		"str": *GetResp(&genre),
	})
}

func GetResp(genre *string) *string {
	address := "localhost:9090"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return &errMsg
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
	return &res.Body
}
