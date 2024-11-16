package main

import (
	"expense-tracker/argparser"
)

type Flag = argparser.Flag

func Add(description string, amount int) {

}

func Update(id int, description string, amount int) {

}

func Delete(id int) {

}

func List(month int) {

}

func Summary(month int) {

}

func main() {
	argparser.AddCmd("add", Add, []Flag{
		{Name: "description", Value: "", Help: "expense description"},
		{Name: "amount", Value: -1, Help: "expense amount"},
	})
	argparser.AddCmd("update", Update, []Flag{
		{Name: "id", Value: -1, Help: "expense id"},
		{Name: "description", Value: "", Help: "expense description"},
		{Name: "amount", Value: -1, Help: "expense amount"},
	})
	argparser.AddCmd("delete", Delete, []Flag{
		{Name: "id", Value: -1, Help: "expense id"},
	})
	argparser.AddCmd("list", List, []Flag{
		{Name: "month", Value: -1, Help: "optional, list only selected month"},
	})
	argparser.AddCmd("summary", Summary, []Flag{
		{Name: "month", Value: -1, Help: "optional, show summaru only for selected month"},
	})

	argparser.Parse()
}
