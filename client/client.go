package main

import (
	"context"
	"fmt"
	"log"

	stockpb "github.com/bradleybonitatibus/stock-rpc/stock"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to dial gRPC server")
	}
	defer con.Close()

	client := stockpb.NewStockServiceClient(con)
	getQuote(client)
}

func getQuote(client stockpb.StockServiceClient) {
	quoteReq := &stockpb.StockQuoteRequest{
		Symbol: "MSFT",
	}
	res, err := client.Quote(context.Background(), quoteReq)

	if err != nil {
		log.Fatal("Failed to fetch quote from Stock Service")
		log.Fatal(err.Error())
	}

	data := res.GetData()
	fmt.Println("Open: ", data.GetOpen())
	fmt.Println("Low: ", data.GetLow())
	fmt.Println("High: ", data.GetHigh())
	fmt.Println("Close: ", data.GetClose())
	fmt.Println("Volume: ", data.GetVolume())
}
