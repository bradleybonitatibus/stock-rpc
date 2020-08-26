package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	stockpb "github.com/bradleybonitatibus/stock-rpc/stock"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type stockServer struct{}

func buildHTTPRequest(method string, symbol string) *http.Request {
	var funcName string
	switch method {
	case "quote":
		funcName = "GLOBAL_QUOTE"
		break
	}
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=%v&symbol=%v&apikey=%v",
		funcName,
		symbol,
		os.Getenv("API_KEY"),
	)
	req, _ := http.NewRequest("GET", url, nil)

	return req
}

func buildHTTPClient() *http.Client {
	client := http.Client{}
	return &client
}

func (*stockServer) Quote(c context.Context, req *stockpb.StockQuoteRequest) (*stockpb.StockQuoteResponse, error) {
	fmt.Println("Called Quote")
	apiReq := buildHTTPRequest("quote", req.GetSymbol())
	client := buildHTTPClient()

	res, err := client.Do(apiReq)

	if err != nil {
		log.Fatal("Failed to call Alphavantage API")
		log.Fatal(err.Error())
		return nil, err
	}
	bytes, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal("Failed to log API Body")
		return nil, err
	}
	fmt.Println("Copied ", bytes, " to stdout")
	return nil, nil
}

func (*stockServer) GetTimeSeriesData(c context.Context, req *stockpb.TimeSeriesRequest) (*stockpb.TimeSeriesResponse, error) {
	return nil, nil
}

func (*stockServer) GetTimeSeriesDataStream(req *stockpb.TimeSeriesRequest, stream stockpb.StockService_GetTimeSeriesDataStreamServer) error {
	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen on port 50051")
		os.Exit(1)
	}

	s := grpc.NewServer()

	stockpb.RegisterStockServiceServer(s, &stockServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to start gRPC serer")
		os.Exit(1)
	}

}
