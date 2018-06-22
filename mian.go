package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	argSize := len(os.Args)
	fmt.Println("input args", args[1:argSize])
	// start

	// end
	fmt.Println()
	fmt.Println("press enter to continue...")
	fmt.Scanln()
}
