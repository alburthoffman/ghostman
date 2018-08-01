package cmd_test

import (
	"fmt"
	"github.com/newly12/ghostman/cmd"
)

func ExampleListCommands() {
	cmds := *(cmd.ListCommands("../examples/templates"))
	fmt.Printf("%v", cmds["get"]["index"])

	// Output:
	// &{{access /} {http 80  / } {} [map[Accept-Language:en-US] map[Cache-Control:no-store]]}
}
