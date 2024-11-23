package expense

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

// helpers

func sort_by_id() []ExpenseID {
	ids := make([]ExpenseID, 0, len(expense_map))
	for k := range expense_map {
		ids = append(ids, k)
	}
	sort.Ints(ids)
	return ids
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

func current_date() ExpenseTime {
	return time.Now()
}

func stringify_month(month int) string {
	return time.Month(month).String()
}

func stringify_current_month() string {
	return stringify_month(int(current_date().Month()))
}

func get_summary(month int, category ExpenseCategory) ExpenseAmount {
	var summary ExpenseAmount
	for _, v := range expense_map {
		if v.filter(month, category) {
			summary += v.Amount
		}
	}
	return summary
}

func check_if_exceeds_budget(amount ExpenseAmount) bool {
	budget, ok := expense_map[budget_id]
	mon := current_date().Month()
	if ok && budget.Date.Month() == mon {
		if amount < 0 {
			amount = get_summary(int(mon), "")
		}
		if amount > budget.Amount {
			fmt.Printf("Monthly expenses exceed monthly budget! Expenses for %s %v, budget %v\n",
				stringify_current_month(), amount, budget.Amount)
			return true
		}
	}
	return false
}

// functions

// add new expense
func Add(description string, amount ExpenseAmount, category ExpenseCategory) {
	if amount <= 0 {
		log.Fatalf("Invalid amount %d", amount)
	}

	last_id++
	expense_map[last_id] = Expense{
		ID:          last_id,
		Description: description,
		Amount:      amount,
		Date:        current_date(),
		Category:    category,
	}

	check_if_exceeds_budget(-1)

	fmt.Printf("Expense added successfully (ID: %d)\n", last_id)
}

// update existing expense
func Update(id int, description string, amount ExpenseAmount, category ExpenseCategory) {
	exp, id := find_expense(id)
	if id < 0 {
		return
	}
	if description != "" {
		exp.Description = description
	}
	if amount > 0 {
		exp.Amount = amount
	}
	if category != "" {
		exp.Category = category
	}
	expense_map[id] = *exp

	check_if_exceeds_budget(-1)

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
func List(month int, category ExpenseCategory) {
	// sort by ids
	ids := sort_by_id()
	fmt.Printf("%-10s %-12s %-20s %-10s %-20s\n",
		expense_header[0], // ID
		expense_header[3], // Date
		expense_header[1], // Descr
		expense_header[2], // Amount
		expense_header[4], // Category
	)

	for _, k := range ids {
		v := expense_map[k]
		print_expense := func() {
			fmt.Printf("%-10v %-12s %-20s %-10v %-20s\n",
				k, v.Date.Format("2006-01-02"), v.Description, v.Amount, v.Category)
		}
		if v.filter(month, category) {
			print_expense()
		}
	}
}

// get expenses summary
func Summary(month int, category ExpenseCategory) {
	summary := get_summary(month, category)
	var month_str string
	if month > 0 && month < 13 {
		month_str = stringify_month(month)
	}
	fmt.Printf("Total expenses for %s %d %s: %v\n", month_str, current_date().Year(), category, summary)
}

// write to csv file
func ExportToCSVFile(filename string, month int, category ExpenseCategory) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
		return
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	// Write header
	err = writer.Write(expense_header)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		return
	}
	// sort by ids
	ids := sort_by_id()
	for _, id := range ids {
		e := expense_map[id]
		if e.filter(month, category) {
			err = writer.Write(e.toCSV())
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", filename, err)
				return
			}
		}
	}
	var month_str string
	if month > 0 && month < 13 {
		month_str = stringify_month(month)
	}
	fmt.Printf("Exported expenses %s %d %s successfully to %s\n", month_str, current_date().Year(), category, filename)
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

// set monthly budget
func SetMonthlyBudget(budget ExpenseAmount) {
	if budget <= 0 {
		log.Fatalf("Invalid amount %d", budget)
	}
	expense_map[budget_id] = Expense{
		ID:          budget_id,
		Amount:      budget,
		Description: "monthly budget",
		Category:    "monthly budget",
		Date:        current_date(),
	}

	check_if_exceeds_budget(-1)

	fmt.Printf("Successfully set monthly budget to %v\n", budget)
}

func ShowMonthlyBudget() {
	budget, ok := expense_map[budget_id]
	if ok && budget.Date.Month() == current_date().Month() {
		fmt.Printf("Budget for %s is %v\n",
			stringify_current_month(), budget.Amount)
	} else {
		fmt.Printf("Budget for %s is not set\n", stringify_current_month())
	}
}
