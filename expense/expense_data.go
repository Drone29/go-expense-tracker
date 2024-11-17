package expense

import (
	"encoding/json"
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
	Date        ExpenseTime   `json:"created-at"`
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
