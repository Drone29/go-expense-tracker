package expense

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ExpenseID = int
type ExpenseTime = time.Time
type ExpenseAmount = int

type Expense struct {
	ID          ExpenseID     `json:"id"`
	Description string        `json:"description"`
	Amount      ExpenseAmount `json:"amount"`
	Date        ExpenseTime   `json:"date"`
}

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

// Convert to json string
func StringifyToJson(tasks []Expense) (string, error) {
	js_bytes, err := json.MarshalIndent(tasks, "", "    ")
	return string(js_bytes), err
}

// Writeto json file
func WriteToJsonFile(filename string, tasks []Expense) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "    ")
	return encoder.Encode(tasks)
}

// Read tasks array from file
func ReadJsonFile(filename string) (tasks []Expense, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	js_bytes, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(js_bytes, &tasks)
	return
}
