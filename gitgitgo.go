package main

import (
	"fmt"
	"os"
	"path/filepath"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/subcommands"
	"phil1436/GitGitGo-CLI/pkg/utils"
	"strings"
)

const VERSION = "0.0.1"

var sbCmds []*cmdtool.Subcommand
var globalFlags *cmdtool.FlagSet = cmdtool.NewFlagSet()

func main() {

	s, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while getting current working directory")
		return
	}
	utils.REPONAME = filepath.Base(s)

	InitCommands()

	ExecuteCommand(os.Args, true)
	logger.Log("")
	logger.Log("by Philipp B.")
	logger.LogSL("Have a nice day :)")
}

func InitCommands() {
	sbCmds = append(sbCmds, cmdtool.NewSubcommand([]string{"version", "-v"}, "Get the current version", PrintVersion))
	sbCmds = append(sbCmds, cmdtool.NewSubcommand([]string{"help", "-h"}, "Help command", globalHelp))

	initCommand := cmdtool.NewSubcommand([]string{"init", "i"}, "Init command", subcommands.Init)

	addCommand := cmdtool.NewSubcommand([]string{"add", "a"}, "Add command", subcommands.Add)

	printCommand := cmdtool.NewSubcommand([]string{"print", "p"}, "Print command", subcommands.Print)

	//shellCommand := cmdtool.NewSubcommand([]string{"shell", "s"}, "Start a GitGitGo shell", subcommands.Print)

	destinationFlag := &cmdtool.Flag{
		Name:        []string{"destination", "d"},
		Description: "<path>: Destination directory",
		Value:       ".",
	}

	fileFlag := &cmdtool.Flag{
		Name:        []string{"file", "f"},
		Description: "<filename>: The file to add. Required for the add command",
		Value:       nil,
	}

	asFlag := &cmdtool.Flag{
		Name:        []string{"as"},
		Description: "<filename>: Give the file a new name",
		Value:       "",
	}

	fileFlagOpt := &cmdtool.Flag{
		Name:        []string{"file", "f"},
		Description: "<filename>: The file to print",
		Value:       "",
	}

	nameFlagPrint := &cmdtool.Flag{
		Name:        []string{"name", "n"},
		Description: ": If enabled only prints the names",
		Value:       false,
		BoolFlag:    true,
	}

	forceFlag := &cmdtool.Flag{
		Name:        []string{"force"},
		Description: ": If enabled files will be overwritten",
		Value:       false,
		BoolFlag:    true,
	}
	dryrunFlag := &cmdtool.Flag{
		Name:        []string{"dry-run", "dryrun", "dr"},
		Description: ": If enabled no files will be created but the output will be printed",
		Value:       false,
		BoolFlag:    true,
	}

	initCommand.AddFlag(destinationFlag)
	initCommand.AddFlag(forceFlag)
	initCommand.AddFlag(dryrunFlag)
	addCommand.AddFlag(fileFlag)
	addCommand.AddFlag(destinationFlag)
	addCommand.AddFlag(asFlag)
	addCommand.AddFlag(forceFlag)
	addCommand.AddFlag(dryrunFlag)
	printCommand.AddFlag(fileFlagOpt)
	printCommand.AddFlag(nameFlagPrint)

	sbCmds = append(sbCmds, initCommand)
	sbCmds = append(sbCmds, addCommand)
	sbCmds = append(sbCmds, printCommand)
	//sbCmds = append(sbCmds, shellCommand)

	// global flags

	quietFlag := &cmdtool.Flag{
		Name:        []string{"quiet", "q"},
		Description: ": If enabled no output is printed",
		Value:       false,
		BoolFlag:    true,
	}

	ownerFlag := &cmdtool.Flag{
		Name:        []string{"owner", "o"},
		Description: "<name>: Change the owner of the .gitgitgo repository",
		Value:       "",
	}
	githubnameFlag := &cmdtool.Flag{
		Name:        []string{"githubname", "gname", "gn"},
		Description: "<name>: Change the github name that will be used",
		Value:       "",
	}
	fullnameFlag := &cmdtool.Flag{
		Name:        []string{"fullname", "fname", "fn"},
		Description: "<name>: Change the full name that will be used",
		Value:       "",
	}
	reponameFlag := &cmdtool.Flag{
		Name:        []string{"reponame", "rname", "rn"},
		Description: "<name>: Change the repo name that will be used",
		Value:       "",
	}

	globalFlags.AddFlag(quietFlag)
	globalFlags.AddFlag(ownerFlag)
	globalFlags.AddFlag(githubnameFlag)
	globalFlags.AddFlag(fullnameFlag)
	globalFlags.AddFlag(reponameFlag)
}

func ExecuteCommand(args []string, ignoreFirstArg bool) {
	activeSCmd := GetActiveSubcommand(args, ignoreFirstArg)

	if activeSCmd == nil {
		globalHelp(nil)
		return
	}
	start := 1
	if ignoreFirstArg {
		start = 2
	}

	parsedGloablFlags := globalFlags.Copy()
	parsedGloablFlags.Parse(args[start:])
	BeforeRun(parsedGloablFlags)

	if !activeSCmd.Run(os.Args[start:], parsedGloablFlags) {
		fmt.Println("\nAn Error occured while executing the command '" + activeSCmd.Names[0] + "':")
		logger.PrintErrors()
		fmt.Println("\nUse 'gitgitgo help' to get help")
	}
}

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
	if !globalFlags.IsEmpty() {
		fmt.Println("Global Flags (Use with every command):")
		fmt.Println(globalFlags.ToString())
	}
	return true
}

func PrintVersion(fs *cmdtool.FlagSet) bool {
	fmt.Printf("GitGitGo-CLI version %s\n", VERSION)
	return true
}

/* func log(msg string) {
	if quietly {
		return
	}
	fmt.Println(msg)
} */
