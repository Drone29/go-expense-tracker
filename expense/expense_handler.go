package expense

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

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
		Date:        time.Now(),
		Category:    category,
	}
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

func check_condition(e Expense, month int, category ExpenseCategory) bool {
	month_filter := month > 0 && month < 13 && int(e.Date.Month()) == month
	category_filter := category != "" && e.Category == category
	return (month <= 0 || month_filter) && (category == "" || category_filter)
}

func sort_by_id() []ExpenseID {
	ids := make([]ExpenseID, 0, len(expense_map))
	for k := range expense_map {
		ids = append(ids, k)
	}
	sort.Ints(ids)
	return ids
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
		if check_condition(v, month, category) {
			print_expense()
		}
	}
}

// get expenses summary
func Summary(month int, category ExpenseCategory) {
	var summary ExpenseAmount
	for _, v := range expense_map {
		if check_condition(v, month, category) {
			summary += v.Amount
		}
	}
	var month_str string
	if month > 0 && month < 13 {
		month_str = time.Month(month).String()
	}
	fmt.Printf("Total expenses for %s %d %s: %v\n", month_str, time.Now().Year(), category, summary)
}

// Write to csv file
func ExportToCSVFile(filename string, month int, category ExpenseCategory) {
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
	// sort by ids
	ids := sort_by_id()
	for _, id := range ids {
		e := expense_map[id]
		if check_condition(e, month, category) {
			err = writer.Write(e.toCSV())
			if err != nil {
				fmt.Printf("Error writing to file %s: %v", filename, err)
				return
			}
		}
	}
	var month_str string
	if month > 0 && month < 13 {
		month_str = time.Month(month).String()
	}
	fmt.Printf("Exported expenses %s %d %s successfully to %s\n", month_str, time.Now().Year(), category, filename)
}
