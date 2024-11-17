package expense

import (
	"fmt"
	"log"
	"time"
)

const expense_json_file = "expenses.json"

var (
	expense_map = map[ExpenseID]Expense{}
	last_id     ExpenseID
)

func find_expense(id int) (*Expense, int) {
	exp, ok := expense_map[id]
	if !ok {
		fmt.Println("No task with id", id)
		return nil, -1
	}
	return &exp, id
}

// load expenses from file
func LoadExpenses() {
	tasks, _ := ReadJsonFile(expense_json_file)
	for _, tsk := range tasks {
		expense_map[tsk.ID] = tsk
		last_id = max(last_id, tsk.ID)
	}
}

// save expenses to file
func SaveExpenses() {
	tasks := make([]Expense, 0, len(expense_map))
	for _, tsk := range expense_map {
		tasks = append(tasks, tsk)
	}
	WriteToJsonFile(expense_json_file, tasks)
}

// add new expense
func Add(description string, amount int) {
	if amount <= 0 {
		log.Fatalf("Invalid amount %d", amount)
	}

	last_id++
	expense_map[last_id] = Expense{
		ID:          last_id,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
	fmt.Printf("Expense added successfully (ID: %d)\n", last_id)
}

// update existing expense
func Update(id int, description string, amount int) {
	if amount <= 0 {
		log.Fatalf("Invalid amount %d", amount)
	}
	exp, id := find_expense(id)
	if id < 0 {
		return
	}
	exp.Description = description
	exp.Amount = amount
	expense_map[id] = *exp
	fmt.Printf("Expense updated successfully (ID: %d)\n", id)
}

// delete expense
func Delete(id int) {
	_, id = find_expense(id)
	if id < 0 {
		return
	}
	delete(expense_map, id)
	fmt.Printf("Expense deleted successfully (ID: %d)\n", id)
}

// list expenses
func List(month int) {
	if month > 0 && month < 13 {

	}
}

// get expenses summary
func Summary(month int) {
	if month > 0 && month < 13 {

	}
}