package expense

import (
	"fmt"
	"log"
	"sort"
	"time"
)

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
	// sort by keys
	keys := make([]ExpenseID, 0, len(expense_map))
	for k := range expense_map {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("%-10s %-12s %-20s %-10s\n", "ID", "Date", "Description", "Amount")
	for _, k := range keys {
		v := expense_map[k]
		print_expense := func() {
			fmt.Printf("%-10v %-12s %-20s %-10v\n",
				k, v.Date.Format("2006-01-02"), v.Description, v.Amount)
		}
		if month > 0 && month < 13 {
			// filter by month (current year)
			if v.Date.Year() == time.Now().Year() && int(v.Date.Month()) == month {
				print_expense()
			}
		} else {
			print_expense()
		}
	}
}

// get expenses summary
func Summary(month int) {
	var summary ExpenseAmount
	var month_filter time.Month
	for _, v := range expense_map {
		if month > 0 && month < 13 {
			month_filter = time.Month(month)
			if v.Date.Year() == time.Now().Year() && v.Date.Month() == month_filter {
				summary += v.Amount
			}
		} else {
			summary += v.Amount
		}
	}

	if month_filter > 0 {
		fmt.Printf("Summary for %s %d: %v\n", month_filter.String(), time.Now().Year(), summary)
	} else {
		fmt.Printf("Summary for %d: %v\n", time.Now().Year(), summary)
	}

}
