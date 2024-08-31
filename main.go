package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		//goland:noinspection ALL
		fmt.Println("Usage: ./kubedlv <command> [<args>...]")
		os.Exit(1)
		return
	}

	config := NewDlvConfig()
	dlvArgs, err := GetDlvArgs(os.Args[1], os.Args[2:], config)
	if err != nil {
		fmt.Println("Failed to parse arguments:", err)
		os.Exit(1)
		return
	}

	fmt.Println("Running dlv with args:", dlvArgs)
	err = RunDlv(dlvArgs)
	if err != nil {
		fmt.Println("Failed to run dlv:", err)
	}
}
