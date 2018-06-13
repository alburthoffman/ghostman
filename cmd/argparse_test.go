package cmd_test

import (
	"fmt"
	"github.com/newly12/ghostman/cmd"
)

func ExampleListCommands() {
	cmds := *(cmd.ListCommands("../examples/templates"))
	fmt.Printf("%v", cmds["get"]["root"])

	// Output:
	// &{{access /} { } {http 80  / } {}}
}
