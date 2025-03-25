package main

import (
	"fmt"
	"os"

	"github.com/naterarmstrong/forkable-library/echo"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: %s <message>", args[0])
	}
	echo.Echo(args[1])
	return nil
}
