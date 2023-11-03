package main

import (
	"GitGitGo-CLI/cmdtool"
	"GitGitGo-CLI/subcommands"
	"fmt"
	"os"
	"strings"
)

const VERSION = "0.0.1"

var sbCmds []*cmdtool.Subcommand
var globalFlags *cmdtool.FlagSet = cmdtool.NewFlagSet()

func main() {

	sbCmds = append(sbCmds, cmdtool.NewSubcommand([]string{"version", "-v"}, "Get the current version", PrintVersion))
	sbCmds = append(sbCmds, cmdtool.NewSubcommand([]string{"help", "-h"}, "Help command", globalHelp))

	initCommand := cmdtool.NewSubcommand([]string{"init", "i"}, "Init command", subcommands.Init)

	addCommand := cmdtool.NewSubcommand([]string{"add", "a"}, "Add command", subcommands.Add)

	destinationFlag := &cmdtool.Flag{
		Name:        []string{"destination", "d"},
		Description: "<path> Destination directory",
		Value:       ".",
	}

	fileFlag := &cmdtool.Flag{
		Name:        []string{"file", "f"},
		Description: "<filename> The file to add. Required for the add command",
		Value:       nil,
	}

	initCommand.AddFlag(destinationFlag)
	addCommand.AddFlag(destinationFlag)
	addCommand.AddFlag(fileFlag)

	sbCmds = append(sbCmds, initCommand)
	sbCmds = append(sbCmds, addCommand)

	activeSCmd := getActiveSubcommand()

	if activeSCmd == nil {
		globalHelp(nil)
		return
	}

	var msg string
	if !activeSCmd.Run(os.Args[2:], globalFlags, &msg) {
		fmt.Println("Could not execute command " + activeSCmd.Names[0] + ":")
		fmt.Println(msg)
	}

}

func getActiveSubcommand() *cmdtool.Subcommand {
	// Get the subcommand name
	if len(os.Args) < 2 {
		return nil
	}

	for _, sb := range sbCmds {
		for _, name := range sb.Names {
			if strings.EqualFold(name, os.Args[1]) {
				return sb
			}
		}
	}
	return nil
}

func globalHelp(fs *cmdtool.FlagSet) {
	fmt.Println("GitGitGo-CLI - A CLI tool for Git written in Go")
	fmt.Println("Usage: gitgitgo [command] [flags]")

	fmt.Println("Commands:")
	for _, sb := range sbCmds {
		fmt.Println(sb.ToString())
	}
}

func PrintVersion(fs *cmdtool.FlagSet) {
	fmt.Printf("GitGitGo-CLI version %s\n\nby Philipp B.", VERSION)
}

/* func log(msg string) {
	if quietly {
		return
	}
	fmt.Println(msg)
} */
