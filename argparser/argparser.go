package argparser

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

type cmdType struct {
	action  reflect.Value
	flagSet *flag.FlagSet
	flags   map[string]int // holds flags' keys and corresponding index
}

type Flag struct {
	Name  string
	Value any // default value
	Help  string
}

func (cmd *cmdType) invoke() {
	// Create arguments to pass to the function
	args := make([]reflect.Value, len(cmd.flags))
	// Visit all flags (even those not set) and populate slice with parsed values
	cmd.flagSet.VisitAll(func(f *flag.Flag) {
		if v, ok := f.Value.(flag.Getter); ok {
			if i, ok := cmd.flags[f.Name]; ok {
				args[i] = reflect.ValueOf(v.Get())
			} else {
				fmt.Printf("Key was not parsed %s\n", f.Name)
				return
			}
		}
	})
	// Invoke the function with the flags' values
	cmd.action.Call(args)
}

var commands = map[string]cmdType{}

// register a new command. action should be a function of any signature
// flags' order and values must correspond to the function's signature
func AddCmd(cmd string, action any, flags []Flag) {
	actVal := reflect.ValueOf(action)
	if actVal.Kind() != reflect.Func {
		fmt.Printf("AddCmd: action must be a function, got %T\n", action)
		return
	}

	fs := flag.NewFlagSet(cmd, flag.ExitOnError)
	flagMap := map[string]int{}
	i := 0
	for _, flag := range flags {
		switch flagVal := flag.Value.(type) {
		case string:
			fs.String(flag.Name, flagVal, flag.Help)
		case bool:
			fs.Bool(flag.Name, flagVal, flag.Help)
		case int:
			fs.Int(flag.Name, flagVal, flag.Help)
		default:
			fmt.Printf("Unsupported type %T\n", flagVal)
			return
		}
		flagMap[flag.Name] = i
		i++
	}

	commands[cmd] = cmdType{
		action:  actVal,
		flagSet: fs,
		flags:   flagMap,
	}
}

func showHelp() {
	fmt.Printf("Usage: %s <CMD>\n", os.Args[0])
	fmt.Printf("List of CMDs:\n")
	for k, v := range commands {
		fmt.Println(k)
		v.flagSet.SetOutput(os.Stdout)
		v.flagSet.PrintDefaults()
	}
}

func Parse() int {
	if len(os.Args) > 1 {
		name := os.Args[1]
		cmd, ok := commands[name]
		if ok {
			args := os.Args[2:]
			err := cmd.flagSet.Parse(args)
			if err != nil || cmd.flagSet.NArg() > 0 {
				fmt.Printf("Error parsing %v\n", cmd.flagSet.Args())
				return -1
			}
			// call function
			cmd.invoke()
		} else {
			fmt.Println("Unknown command", name)
			showHelp()
			return -1
		}
	} else {
		showHelp()
	}
	return 0
}
