package argparser

import (
	"os"
	"testing"
)

func TestParserNoDefault(t *testing.T) {
	var test_str string
	test_f := func(val string) {
		test_str = val
	}
	AddCmd("test", test_f, []Flag{
		{Name: "str", Value: ""},
	})
	os.Args = []string{os.Args[0], "test", "--str", "abcd"}
	res := Parse()
	if res != 0 || test_str != "abcd" {
		t.Error("Invalid parse")
	}
}

func TestParserDefault(t *testing.T) {
	var test_str string
	test_f := func(val string) {
		test_str = val
	}
	AddCmd("test", test_f, []Flag{
		{Name: "str", Value: "default string"},
	})
	os.Args = []string{os.Args[0], "test"}
	res := Parse()
	if res != 0 || test_str != "default string" {
		t.Error("Invalid parse")
	}
}

func TestParserOrder(t *testing.T) {
	var test_str string
	var test_int int
	test_f := func(val string, val2 int) {
		test_str = val
		test_int = val2
	}
	AddCmd("test", test_f, []Flag{
		{Name: "str", Value: ""},
		{Name: "int", Value: 0},
	})
	// should parse params and pass them into function no matter in what order they're supplied
	os.Args = []string{os.Args[0], "test", "--int", "456", "--str", "abcd"}
	res := Parse()
	if res != 0 || test_str != "abcd" || test_int != 456 {
		t.Error("Invalid parse")
	}
}

func TestInvalid(t *testing.T) {
	test_f := func(val string, val2 int) {}
	AddCmd("test", test_f, []Flag{
		{Name: "str", Value: ""},
		{Name: "int", Value: 0},
	})
	// invalid args (omitted --int)
	os.Args = []string{os.Args[0], "test", "456", "--str", "abcd"}
	res := Parse()
	if res == 0 {
		t.Error("Should've returned error")
	}
}
