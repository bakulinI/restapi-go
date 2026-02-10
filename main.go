package main

import (
	"fmt"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculations []Calculation{}


func main() {
	fmt.Println("Hello World")
}
