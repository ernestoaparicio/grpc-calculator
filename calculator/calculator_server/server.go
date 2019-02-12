package main

import (
	"context"
	"fmt"
	"github.com/ernestoaparicio/grpc-calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (server) CalculatorAdd(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("CalculatorAdd function was invoked with: %v", req)

	firstNumber := req.GetInputNumbers().GetNumber_1()
	secondNumber := req.GetInputNumbers().GetNumber_2()

	result := firstNumber + secondNumber

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

func (server) CalculatorSubtract(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("CalculatorAdd function was invoked with: %v", req)

	firstNumber := req.GetInputNumbers().GetNumber_1()
	secondNumber := req.GetInputNumbers().GetNumber_2()

	result := firstNumber - secondNumber

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

func (server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculator function was invoked with: %v", req)

	firstNumber := req.GetInputNumbers().GetNumber_1()
	secondNumber := req.GetInputNumbers().GetNumber_2()

	result := firstNumber * secondNumber

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

func main(){
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil{
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	err2 := s.Serve(lis)

	if err2 != nil{
		log.Fatalf("Failed to serve: %v", err2)
	}
}