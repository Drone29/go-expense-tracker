package expense

import (
	"os"
	"testing"
)

func TestWriteRead(t *testing.T) {
	var sample_task = Expense{
		ID:          12,
		Description: "Test",
		Amount:      23,
		Category:    "test",
	}
	if err := writeToJsonFile("test.json", []Expense{sample_task}); err != nil {
		t.Errorf("Error writing to file %v", sample_task)
	}
	new_tasks, err := readJsonFile("test.json")
	if err != nil {
		t.Errorf("Error reading from file %v", err)
	}
	if new_tasks[0] != sample_task {
		t.Errorf("Results do not match! old: %v new: %v", sample_task, new_tasks[0])
	}
	os.Remove("test.json")
}

func TestAddExpense(t *testing.T) {
	for i := 1; i < 11; i++ {
		Add("test", ExpenseAmount(i), "test")
	}

	sum := ExpenseAmount(0)
	for _, v := range expense_map {
		sum += v.Amount
	}

	if sum != ExpenseAmount(55) {
		t.Errorf("Invalid sum! %v", sum)
	}
}

func TestUpdateExpense(t *testing.T) {
	// add expense
	Add("test", 10, "test")
	// update amount
	Update(last_id, "updated", 14, "test")
	// check
	if expense_map[last_id].Amount != 14 || expense_map[last_id].Description != "updated" {
		t.Errorf("Expense was not updated! %v", expense_map[last_id])
	}
}

func TestDeleteExpense(t *testing.T) {
	old_len := len(expense_map)
	Add("test", 10, "test")
	Delete(last_id)
	if len(expense_map) != old_len {
		t.Errorf("Error deleting expense, len=%v", len(expense_map))
	}
}

func TestExpenseFilter(t *testing.T) {
	e := Expense{
		ID:       1,
		Date:     current_date(),
		Category: "test",
	}

	if ok := e.filter(int(current_date().Month()), "test"); !ok {
		t.Error("Should be valid!")
	}
	if ok := e.filter(int(current_date().Month()), ""); !ok {
		t.Error("Should be valid!")
	}
	if ok := e.filter(0, "test"); !ok {
		t.Error("Should be valid!")
	}
	if ok := e.filter(int(current_date().Month()), "tttt"); ok {
		t.Error("Should be invalid!")
	}
	next_month := int(current_date().Month()+1)%12 + 1
	if ok := e.filter(next_month, "test"); ok {
		t.Errorf("Should not have found month %v!", next_month)
	}
	e.ID = budget_id
	if ok := e.filter(int(current_date().Month()), "test"); ok {
		t.Error("Should be invalid!")
	}
}

func TestFindExpense(t *testing.T) {
	expense_map[budget_id] = Expense{
		ID:       budget_id,
		Date:     current_date(),
		Category: "test",
	}
	expense_map[100] = Expense{
		ID:       100,
		Date:     current_date(),
		Category: "test",
	}
	if exp, _ := find_expense(budget_id); exp != nil {
		t.Errorf("Budget object should not be found")
	}
	if exp, _ := find_expense(100); exp == nil {
		t.Errorf("Should've found existing non-budget object")
	}
}
