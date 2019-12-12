package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	api "github.com/atreyagaurav/quotes/pkg/api"
)

type quotesServer struct {
	name string
}

func NewQuotesServer(st *string) api.QuotesServer {
	return &quotesServer{name: *st}
}

func ReadCsv(fileName string, category string) (map[int]string, error) {
	// Open the file
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Print("Couldn't open the csv file", err)
		return nil, err
	}
	m := make(map[int]string, 1)
	r := csv.NewReader(csvfile)
	i := 0
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error in reading", err)
			return nil, err
		}
		if category == record[2] {
			m[i] = fmt.Sprintf("\"%s\" \n\t -%s\n", record[0], record[1])
			i = i + 1
		}
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}
	return m, nil
}

//temp function to make the csv read/write
// func main() {
// 	records, err := ReadCsv("../../data/quotes.csv", "love")
// 	if err != nil {
// 		log.Print("Error ", err)
// 	}
// 	fmt.Println(records)
// }

func (q *quotesServer) Get(ctx context.Context, req *api.QuoteRequest) (*api.QuoteResponse, error) {

	records, err := ReadCsv("../../data/quotes.csv", req.Genre)
	if err != nil {
		log.Print("Error ", err)
		return nil, err
	}
	// fmt.Println(records)
	quote := "There is no greater impediment to the advancement of knowledge than the ambiguity of words.\n\t -Unknown\n"
	if len(records) > 0 {
		index := rand.Intn(len(records))
		quote = records[index]
	}
	return &api.QuoteResponse{
		Id:     req.Id,
		Genre:  req.Genre,
		Body:   quote,
		Author: "gaurav",
	}, nil
}
