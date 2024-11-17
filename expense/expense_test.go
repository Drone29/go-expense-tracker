package expense

import (
	"encoding/json"
	"os"
	"testing"
)

var sample_task = Expense{
	ID:          12,
	Description: "Test",
}

func TestStringify(t *testing.T) {
	taskstr, err := StringifyToJson([]Expense{sample_task})
	if err != nil {
		t.Errorf("Error stringifying task [%v] %v", sample_task, err)
	}
	var new_tasks []Expense
	if err = json.Unmarshal([]byte(taskstr), &new_tasks); err != nil {
		t.Errorf("Error unmarshaling string %v %v", taskstr, err)
	}
	if new_tasks[0] != sample_task {
		t.Errorf("Results do not match! old: %v new: %v", sample_task, new_tasks[0])
	}
}

func TestWriteRead(t *testing.T) {
	if err := WriteToJsonFile("test.json", []Expense{sample_task}); err != nil {
		t.Errorf("Error writing to file %v", sample_task)
	}
	new_tasks, err := ReadJsonFile("test.json")
	if err != nil {
		t.Errorf("Error reading from file %v", err)
	}
	if new_tasks[0] != sample_task {
		t.Errorf("Results do not match! old: %v new: %v", sample_task, new_tasks[0])
	}
	os.Remove("test.json")
}
