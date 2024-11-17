package main

import (
	"expense-tracker/argparser"
	"expense-tracker/expense"
	"fmt"
	"time"
)

type Flag = argparser.Flag

func main() {

	expense.LoadExpenses()

	argparser.AddCmd("add", expense.Add, []Flag{
		{Name: "description", Value: "", Help: "expense description"},
		{Name: "amount", Value: -1, Help: "expense amount"},
		{Name: "category", Value: "", Help: "optional, expense category"},
	})
	argparser.AddCmd("update", expense.Update, []Flag{
		{Name: "id", Value: -1, Help: "expense id"},
		{Name: "description", Value: "", Help: "expense description"},
		{Name: "amount", Value: -1, Help: "expense amount"},
		{Name: "category", Value: "", Help: "optional, expense category"},
	})
	argparser.AddCmd("delete", expense.Delete, []Flag{
		{Name: "id", Value: -1, Help: "expense id"},
	})
	argparser.AddCmd("list", expense.List, []Flag{
		{Name: "month", Value: -1, Help: "optional, list only selected month"},
		{Name: "category", Value: "", Help: "optional, expense category"},
	})
	argparser.AddCmd("summary", expense.Summary, []Flag{
		{Name: "month", Value: -1, Help: "optional, show summary only for selected month"},
		{Name: "category", Value: "", Help: "optional, expense category"},
	})
	argparser.AddCmd("export-csv", expense.ExportToCSVFile, []Flag{
		{Name: "filename", Value: fmt.Sprintf("expenses-%d.csv", time.Now().Year())},
	})

	argparser.Parse()

	expense.SaveExpenses()
}
