package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// check if the user provided a subcommand
	if len(os.Args) < 2 {
		fmt.Println("install subcommand is required")
		os.Exit(1)
	}

	// create a new flag set for the install subcommand
	installCommand := flag.NewFlagSet("install", flag.ExitOnError)
	testCommand := flag.NewFlagSet("test", flag.ExitOnError)
	fmt.Println(installCommand.Args())
	fmt.Println(testCommand.Args())
	// parse the flags for the install flag set
	installCommand.Parse(os.Args[2:])
	testCommand.Parse(os.Args[2:])
	// check if the install flag set was parsed
	if installCommand.Parsed() {
		hallo()
	}
	if testCommand.Parsed() {
		hallo2()
	}

}

func hallo() {
	fmt.Println("install!")
}
func hallo2() {
	fmt.Println("test!")
}
