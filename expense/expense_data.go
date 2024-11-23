package expense

import (
	"fmt"
	"time"
)

const budget_id ExpenseID = 0     // special id for budget object
const DefAmt ExpenseAmount = -1.0 // default amount

type ExpenseID = int
type ExpenseTime = time.Time
type ExpenseAmount = float64
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
