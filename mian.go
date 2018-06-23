package main

import (
	"fmt"
	"os"
	"sys/file"
	"sys/timetool"
	"time"
)

func main() {
	args := os.Args
	argSize := len(os.Args)
	// fmt.Println("input args", args[1:argSize])
	// start
	now := time.Now()
	for i := 1; i < argSize; i++ {
		itFile := file.NewFile(args[i])
		afterTime := now.Add(time.Second * time.Duration(i))
		itFile.Name = timetool.GetTime(afterTime)
		file.Rename(itFile)
		fmt.Println(itFile.GetPath())
	}

	// end
	fmt.Println()
	fmt.Println("press enter to continue...")
	fmt.Scanln()
}
