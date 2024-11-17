package expense

import (
	"encoding/csv"
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

func (e *Expense) toCSV() []string {
	return []string{
		fmt.Sprintf("%v", e.ID),     // ID
		e.Description,               // Descr
		fmt.Sprintf("%v", e.Amount), // Amount
		e.Date.Format("2006-01-02"), // Date
	}
}

var (
	expense_map    = map[ExpenseID]Expense{}
	last_id        ExpenseID
	expense_header = []string{
		"ID", "Description", "Amount", "Date",
	}
)

const expense_json_file = "expenses.json"

func find_expense(id int) (*Expense, int) {
	exp, ok := expense_map[id]
	if !ok {
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

// Write to csv file
func WriteToCSVFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v", filename, err)
		return
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	// Write header
	err = writer.Write(expense_header)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v", filename, err)
		return
	}
	for _, e := range expense_map {
		err = writer.Write(e.toCSV())
		if err != nil {
			fmt.Printf("Error writing to file %s: %v", filename, err)
			return
		}
	}
}

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

// Convert to json string
func StringifyToJson(tasks []Expense) (string, error) {
	js_bytes, err := json.MarshalIndent(tasks, "", "    ")
	return string(js_bytes), err
}
