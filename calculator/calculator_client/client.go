package main

import (
	"context"
	"fmt"
	"github.com/ernestoaparicio/grpc-calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Println("Client running.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Unable to connect to: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting Unary RPC")

	req := &calculatorpb.CalculatorRequest{
		InputNumbers: &calculatorpb.InputNumbers{
			Number_1:2,
			Number_2:2,
		},
	}

	res, err := c.Calculator(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while calling Calculator RPC: %v", err)
	}

	log.Printf("Response from calculator: %v", res.Result)
}