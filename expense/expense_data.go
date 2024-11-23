package expense

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// special id for budget object
const budget_id ExpenseID = 0

type ExpenseID = int
type ExpenseTime = time.Time
type ExpenseAmount = int
type ExpenseCategory = string

type Expense struct {
	ID          ExpenseID       `json:"id"`
	Description string          `json:"description"`
	Amount      ExpenseAmount   `json:"amount"`
	Date        ExpenseTime     `json:"date"`
	Category    ExpenseCategory `json:"category,omitempty"`
}

var (
	expense_map    = map[ExpenseID]Expense{}
	last_id        ExpenseID
	expense_header = []string{
		"ID", "Description", "Amount", "Date", "Category",
	}
	expense_json_file = "expenses.json"
)

// Expense methods
func (e *Expense) toCSV() []string {
	return []string{
		fmt.Sprintf("%v", e.ID),     // ID
		e.Description,               // Descr
		fmt.Sprintf("%v", e.Amount), // Amount
		e.Date.Format("2006-01-02"), // Date
		e.Category,
	}
}

func (e *Expense) filter(month int, category ExpenseCategory) bool {
	month_filter := month > 0 && month < 13 && int(e.Date.Month()) == month
	category_filter := category != "" && e.Category == category
	id_cond := e.ID != budget_id
	mon_cond := month <= 0 || month_filter
	cat_cond := category == "" || category_filter
	return id_cond && mon_cond && cat_cond
}

// helpers

func current_date() ExpenseTime {
	return time.Now()
}

func find_expense(id int) (*Expense, int) {
	exp, ok := expense_map[id]
	if !ok || id == budget_id {
		fmt.Println("No task with id", id)
		return nil, -1
	}
	return &exp, id
}

// Write to json file
func writeToJsonFile(filename string, expenses []Expense) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "    ")
	return encoder.Encode(expenses)
}

// Read tasks array from file
func readJsonFile(filename string) (expenses []Expense, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	js_bytes, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(js_bytes, &expenses)
	return
}

// functions
// load expenses from file
func LoadExpenses() {
	tasks, _ := readJsonFile(expense_json_file)
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
	writeToJsonFile(expense_json_file, tasks)
}
