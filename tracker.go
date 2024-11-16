package main

import (
	"expense-tracker/argparser"
	"fmt"
)

type Flag = argparser.Flag

func Add(description string, amount int) {
	fmt.Printf("Test %s %d", description, amount)
}

func main() {
	argparser.AddCmd("add", Add, []Flag{
		{Name: "description", Value: "", Help: "expense description"},
		{Name: "amount", Value: 0, Help: "expense amount"},
	})

	argparser.Parse()
}
