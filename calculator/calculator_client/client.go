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

	doAdd(c)
	doSub(c)
}

func doAdd(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting Unary CalculatorAdd RPC")

	req := &calculatorpb.CalculatorRequest{
		InputNumbers: &calculatorpb.InputNumbers{
			Number_1:2,
			Number_2:2,
		},
	}

	res, err := c.CalculatorAdd(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while calling CalculatorAdd RPC: %v", err)
	}

	log.Printf("Response from calculator: %v", res.Result)
}

func doSub(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting Unary CalculatorSub RPC")

	req := &calculatorpb.CalculatorRequest{
		InputNumbers: &calculatorpb.InputNumbers{
			Number_1:2,
			Number_2:2,
		},
	}

	res, err := c.CalculatorSubtract(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while calling CalculatorSub RPC: %v", err)
	}

	log.Printf("Response from calculator: %v", res.Result)
}