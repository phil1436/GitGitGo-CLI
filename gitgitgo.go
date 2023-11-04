package main

import (
	"fmt"
	"os"
	"path/filepath"
	"phil1436/GitGitGo-CLI/cmdtool"
	"phil1436/GitGitGo-CLI/subcommands"
	"phil1436/GitGitGo-CLI/utils"
	"strings"
)

const VERSION = "0.0.1"

var sbCmds []*cmdtool.Subcommand
var globalFlags *cmdtool.FlagSet = cmdtool.NewFlagSet()

func main() {

	//
	s, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while getting current working directory")
		return
	}

	utils.REPONAME = filepath.Base(s)

	sbCmds = append(sbCmds, cmdtool.NewSubcommand([]string{"version", "-v"}, "Get the current version", PrintVersion))
	sbCmds = append(sbCmds, cmdtool.NewSubcommand([]string{"help", "-h"}, "Help command", globalHelp))

	initCommand := cmdtool.NewSubcommand([]string{"init", "i"}, "Init command", subcommands.Init)

	addCommand := cmdtool.NewSubcommand([]string{"add", "a"}, "Add command", subcommands.Add)

	printCommand := cmdtool.NewSubcommand([]string{"print", "p"}, "Print command", subcommands.Print)

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

	fileFlagOpt := &cmdtool.Flag{
		Name:        []string{"file", "f"},
		Description: "<filename> The file to print",
		Value:       "",
	}

	nameFlagPrint := &cmdtool.Flag{
		Name:        []string{"name", "n"},
		Description: "If enabled only prints the name",
		Value:       false,
		BoolFlag:    true,
	}

	initCommand.AddFlag(destinationFlag)
	addCommand.AddFlag(destinationFlag)
	addCommand.AddFlag(fileFlag)
	printCommand.AddFlag(fileFlagOpt)
	printCommand.AddFlag(nameFlagPrint)

	sbCmds = append(sbCmds, initCommand)
	sbCmds = append(sbCmds, addCommand)
	sbCmds = append(sbCmds, printCommand)

	activeSCmd := getActiveSubcommand()

	if activeSCmd == nil {
		globalHelp(nil)
		return
	}

	var msg string
	if !activeSCmd.Run(os.Args[2:], globalFlags, &msg) {
		fmt.Println("Error while executing command " + activeSCmd.Names[0] + ":")
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

func globalHelp(fs *cmdtool.FlagSet) bool {
	fmt.Println("GitGitGo-CLI - A CLI tool for Git written in Go")
	fmt.Println("")
	//fmt.Printf("Usage: %s [command] [flags]\n", os.Args[0])
	fmt.Println("Usage: gitgitgo [command] [flags]")
	fmt.Println("")
	fmt.Println("Commands:")
	for _, sb := range sbCmds {
		fmt.Println(sb.ToString())
	}
	fmt.Println("")
	fmt.Println("by Philipp B.")
	return true
}

func PrintVersion(fs *cmdtool.FlagSet) bool {
	fmt.Printf("GitGitGo-CLI version %s\n\nby Philipp B.", VERSION)
	return true
}

/* func log(msg string) {
	if quietly {
		return
	}
	fmt.Println(msg)
} */
