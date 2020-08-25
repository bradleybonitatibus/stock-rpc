package main

import (
	"context"
	"log"
	"net"
	"os"

	stockpb "github.com/bradleybonitatibus/stock-rpc/stock"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type stockServer struct{}

func (*stockServer) Quote(c context.Context, req *stockpb.StockQuoteRequest) (*stockpb.StockQuoteResponse, error) {
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
