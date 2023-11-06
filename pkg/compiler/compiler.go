package compiler

import (
	"fmt"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/utils"
	"strings"
)

const VERSION = "0.0.1"

var sbCmds []*cmdtool.Subcommand
var globalFlags *cmdtool.FlagSet = cmdtool.NewFlagSet()

// ExecuteCommand executes the command
func ExecuteCommand(args []string, ignoreFirstArg bool) {
	activeSCmd := GetActiveSubcommand(args, ignoreFirstArg)

	if activeSCmd == nil {
		GlobalHelp(nil)
		return
	}
	start := 1
	if ignoreFirstArg {
		start = 2
	}

	parsedGloablFlags := globalFlags.Copy()
	parsedGloablFlags.Parse(args[start:])
	BeforeRun(parsedGloablFlags)

	if !activeSCmd.Run(args[start:], parsedGloablFlags) {
		fmt.Println("\nAn Error occured while executing the command '" + activeSCmd.Names[0] + "':")
		logger.PrintErrors()
		fmt.Println("\nUse 'gitgitgo help' to get help")
	}
}

// BeforeRun is called before the subcommand is executed
func BeforeRun(globalFlags *cmdtool.FlagSet) {
	if globalFlags.GetValue("quiet").(bool) {
		logger.Quiet()
	}
	if globalFlags.GetValue("owner").(string) != "" {
		utils.OWNER = globalFlags.GetValue("owner").(string)
	}
	if globalFlags.GetValue("githubname").(string) != "" {
		utils.GITHUBNAME = globalFlags.GetValue("githubname").(string)
	}
	if globalFlags.GetValue("fullname").(string) != "" {
		utils.FULLNAME = globalFlags.GetValue("fullname").(string)
	}
	if globalFlags.GetValue("reponame").(string) != "" {
		utils.REPONAME = globalFlags.GetValue("reponame").(string)
	}
}

// GetActiveSubcommand returns the active subcommand
func GetActiveSubcommand(args []string, ignoreFirstArg bool) *cmdtool.Subcommand {
	// Get the subcommand name
	first := 0

	if ignoreFirstArg {
		first = 1
	}

	if len(args) < (first + 1) {
		return nil
	}

	for _, sb := range sbCmds {
		for _, name := range sb.Names {
			if strings.EqualFold(name, args[first]) {
				return sb
			}
		}
	}
	return nil
}

// GlobalHelp prints the global help
func GlobalHelp(fs *cmdtool.FlagSet) bool {
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
	if !globalFlags.IsEmpty() {
		fmt.Println("Global Flags (Use with every command):")
		fmt.Println(globalFlags.ToString())
	}
	return true
}

// PrintVersion prints the current version
func PrintVersion(fs *cmdtool.FlagSet) bool {
	fmt.Printf("GitGitGo-CLI version %s\n", VERSION)
	return true
}

func AddSubcommand(scmd *cmdtool.Subcommand) {
	sbCmds = append(sbCmds, scmd)
}

func AddGlobalFlag(flag *cmdtool.Flag) {
	globalFlags.AddFlag(flag)
}
