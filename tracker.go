package main

import (
	"expense-tracker/argparser"
	"fmt"
)

func Add(description string, amount int) {
	fmt.Printf("Test %s %d", description, amount)
}

func main() {
	argparser.AddCmd("add", Add, []argparser.Flag{
		{Name: "description", Value: ""},
		{Name: "amount", Value: 0},
	})

	argparser.Parse()
}
