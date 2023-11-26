package main

import (
	"os"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/compiler"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/subcommands"
	"phil1436/GitGitGo-CLI/pkg/utils"
)

func main() {

	utils.SetRepoName()

	InitCommands()

	compiler.ExecuteCommand(os.Args, true)

	logger.Log("")
	logger.Log("by Philipp B.")
	logger.Log("Have a nice day :)")
}

func InitCommands() {
	compiler.AddSubcommand(cmdtool.NewSubcommand([]string{"version", "-v"}, "Get the current version", compiler.PrintVersion))
	helpCommad := cmdtool.NewSubcommand([]string{"help", "-h"}, "Help command", compiler.Help)
	helpCommad.AddAttribute("command", "The command to get help for")
	compiler.AddSubcommand(helpCommad)

	initCommand := cmdtool.NewSubcommand([]string{"init", "i"}, "Init command", subcommands.Init)

	initCommand.AddAttribute("keywords", "A comma separated list of keywords to search for")

	addCommand := cmdtool.NewSubcommand([]string{"add", "a"}, "Add command", subcommands.Add)

	addCommand.AddAttribute("filename", "The file to add")

	printCommand := cmdtool.NewSubcommand([]string{"print", "p"}, "Print command", subcommands.Print)

	printCommand.AddAttribute("filename | parametername", "The file or value to print (If no file is given all files will be printed)")

	shellCommand := cmdtool.NewSubcommand([]string{"shell", "s"}, "Start a GitGitGo shell", subcommands.StartShell)

	runCommand := cmdtool.NewSubcommand([]string{"run", "r"}, "Run a specified file", subcommands.Run)

	runCommand.AddAttribute("filename", "The file to run")

	setCommand := cmdtool.NewSubcommand([]string{"set"}, "Set a parameter", func(attValue []interface{}, fs *cmdtool.FlagSet) bool {

		if attValue[0] == nil {
			logger.AddError("No name specified")
			return false
		}
		if attValue[1] == nil {
			logger.AddError("No value specified")
			return false
		}

		name := attValue[0].(string)
		value := attValue[1].(string)

		if utils.IsRepoVarName(name) {
			utils.REPONAME = value
			name = "reponame"
		} else if utils.IsProviderVarName(name) {
			utils.PROVIDER = value
			name = "provider"
		} else if utils.IsFullNameVarName(name) {
			utils.FULLNAME = value
			name = "fullname"
		} else if utils.IsGithubNameVarName(name) {
			utils.GITHUBNAME = value
			name = "githubname"
		} else {
			logger.AddError("Unknown name: " + name)
			return false
		}
		logger.Log("Set " + name + " to " + value)
		return true
	})
	setCommand.AddAttribute("name", "The name of the parameter to set")
	setCommand.AddAttribute("value", "The value to set the parameter to")

	destinationFlag := &cmdtool.Flag{
		Name:        []string{"destination", "d"},
		Description: "<path>: Destination directory",
		Value:       ".",
	}
	asFlag := &cmdtool.Flag{
		Name:        []string{"as"},
		Description: "<filename>: Give the file a new name",
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
	forceRunFlag := &cmdtool.Flag{
		Name:        []string{"force"},
		Description: ": If enabled the execution will continue even if an error occurs",
		Value:       false,
		BoolFlag:    true,
	}
	dryrunFlag := &cmdtool.Flag{
		Name:        []string{"dry-run", "dryrun", "dr"},
		Description: ": If enabled no files will be created but the output will be printed",
		Value:       false,
		BoolFlag:    true,
	}
	ignoreFlag := &cmdtool.Flag{
		Name:        []string{"ignore", "i"},
		Description: "<file1,file2, ... >: Commas separated list of files to ignore",
		Value:       "",
	}

	initCommand.AddFlag(destinationFlag)
	initCommand.AddFlag(forceFlag)
	initCommand.AddFlag(dryrunFlag)
	initCommand.AddFlag(ignoreFlag)
	addCommand.AddFlag(destinationFlag)
	addCommand.AddFlag(asFlag)
	addCommand.AddFlag(forceFlag)
	addCommand.AddFlag(dryrunFlag)
	printCommand.AddFlag(nameFlagPrint)
	runCommand.AddFlag(forceRunFlag)

	compiler.AddSubcommand(initCommand)
	compiler.AddSubcommand(addCommand)
	compiler.AddSubcommand(printCommand)
	compiler.AddSubcommand(shellCommand)
	compiler.AddSubcommand(runCommand)
	compiler.AddSubcommand(setCommand)

	// global flags
	quietFlag := &cmdtool.Flag{
		Name:        []string{"quiet", "q"},
		Description: ": If enabled no output is printed",
		Value:       false,
		BoolFlag:    true,
	}

	providerFlag := &cmdtool.Flag{
		Name:        []string{"provider", "p"},
		Description: "<name>: Change the provider of the .gitgitgo repository",
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
	devFlag := &cmdtool.Flag{
		Name:        []string{"dev"},
		Description: ": Run in dev mode [internal use only]",
		Value:       false,
		BoolFlag:    true,
	}

	compiler.AddGlobalFlag(quietFlag)
	compiler.AddGlobalFlag(providerFlag)
	compiler.AddGlobalFlag(githubnameFlag)
	compiler.AddGlobalFlag(fullnameFlag)
	compiler.AddGlobalFlag(reponameFlag)
	compiler.AddGlobalFlag(devFlag)

}
