package expense

import (
	"os"
	"testing"
)

var sample_task = Expense{
	ID:          12,
	Description: "Test",
	Amount:      23,
	Category:    "test",
}

func TestWriteRead(t *testing.T) {
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
		Add("test", i, "test")
	}

	sum := 0
	for _, v := range expense_map {
		sum += v.Amount
	}

	if sum != 55 {
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
