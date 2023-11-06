package main

import (
	"fmt"
	"os"
	"path/filepath"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/compiler"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/subcommands"
	"phil1436/GitGitGo-CLI/pkg/utils"
)

func main() {

	s, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while getting current working directory")
		return
	}
	utils.REPONAME = filepath.Base(s)
	InitCommands()
	compiler.ExecuteCommand(os.Args, true)

	logger.Log("")
	logger.Log("by Philipp B.")
	logger.LogSL("Have a nice day :)")
}

func InitCommands() {
	compiler.AddSubcommand(cmdtool.NewSubcommand([]string{"version", "-v"}, "Get the current version", compiler.PrintVersion))
	compiler.AddSubcommand(cmdtool.NewSubcommand([]string{"help", "-h"}, "Help command", compiler.GlobalHelp))

	initCommand := cmdtool.NewSubcommand([]string{"init", "i"}, "Init command", subcommands.Init)

	addCommand := cmdtool.NewSubcommand([]string{"add", "a"}, "Add command", subcommands.Add)

	printCommand := cmdtool.NewSubcommand([]string{"print", "p"}, "Print command", subcommands.Print)

	shellCommand := cmdtool.NewSubcommand([]string{"shell", "s"}, "Start a GitGitGo shell", subcommands.StartShell)

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

	compiler.AddSubcommand(initCommand)
	compiler.AddSubcommand(addCommand)
	compiler.AddSubcommand(printCommand)
	compiler.AddSubcommand(shellCommand)

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

	compiler.AddGlobalFlag(quietFlag)
	compiler.AddGlobalFlag(ownerFlag)
	compiler.AddGlobalFlag(githubnameFlag)
	compiler.AddGlobalFlag(fullnameFlag)
	compiler.AddGlobalFlag(reponameFlag)

}
