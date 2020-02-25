package main

import (
	"context"
	"fmt"
	"log"
	"net"

	calculatorpb "github.com/soumitrachakravarti/g-rpc-tutorial/calculator/calculator-pb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// Logging to let us know that server is running
	fmt.Println("Server running...")

	// Creating a listner
	l, e := net.Listen("tcp", "localhost:50051")
	if e != nil {
		log.Fatalf("Couldn't start listener %v\n", e)
	}

	// Creating a server using gRPC
	s := grpc.NewServer()

	// Registering with Calculator Service created for us from the proto file
	calculatorpb.RegisterSumServiceServer(s, &server{})
	calculatorpb.RegisterSubServiceServer(s, &server{})
	calculatorpb.RegisterMulServiceServer(s, &server{})
	calculatorpb.RegisterDivServiceServer(s, &server{})

	if e := s.Serve(l); e != nil {
		log.Fatalf("Failed to serve: %f", e)
	}
}

// Requst Sum function
func (*server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Printf("Sum function invoked by %v\n", request)

	a := request.GetA()
	b := request.GetB()

	// Add numbers
	result := a + b

	response := &calculatorpb.SumResponse{
		Result: result,
	}

	return response, nil
}

// Requst Sub function
func (*server) Sub(ctx context.Context, request *calculatorpb.SubRequest) (*calculatorpb.SubResponse, error) {

	fmt.Printf("Sub function invoked by %v\n", request)

	a := request.GetA()
	b := request.GetB()

	// Subtract numbers
	result := a - b

	response := &calculatorpb.SubResponse{
		Result: result,
	}

	return response, nil
}

// Requst Mul function
func (*server) Mul(ctx context.Context, request *calculatorpb.MulRequest) (*calculatorpb.MulResponse, error) {

	fmt.Printf("Mul function invoked by %v\n", request)

	a := request.GetA()
	b := request.GetB()

	// Multiply numbers
	result := a * b

	response := &calculatorpb.MulResponse{
		Result: result,
	}

	return response, nil
}

// Requst Div function
func (*server) Div(ctx context.Context, request *calculatorpb.DivRequest) (*calculatorpb.DivResponse, error) {

	fmt.Printf("Div function invoked by %v\n", request)

	a := request.GetA()
	b := request.GetB()

	// Div numbers
	result := a / b

	response := &calculatorpb.DivResponse{
		Result: result,
	}

	return response, nil
}
