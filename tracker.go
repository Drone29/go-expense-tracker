package main

import (
	"expense-tracker/argparser"
	"expense-tracker/expense"
)

type Flag = argparser.Flag

func main() {

	expense.LoadExpenses()

	argparser.AddCmd("add", expense.Add, []Flag{
		{Name: "description", Value: "", Help: "expense description"},
		{Name: "amount", Value: -1, Help: "expense amount"},
	})
	argparser.AddCmd("update", expense.Update, []Flag{
		{Name: "id", Value: -1, Help: "expense id"},
		{Name: "description", Value: "", Help: "expense description"},
		{Name: "amount", Value: -1, Help: "expense amount"},
	})
	argparser.AddCmd("delete", expense.Delete, []Flag{
		{Name: "id", Value: -1, Help: "expense id"},
	})
	argparser.AddCmd("list", expense.List, []Flag{
		{Name: "month", Value: -1, Help: "optional, list only selected month"},
	})
	argparser.AddCmd("summary", expense.Summary, []Flag{
		{Name: "month", Value: -1, Help: "optional, show summary only for selected month"},
	})

	argparser.Parse()

	expense.SaveExpenses()
}
