package shell

import (
	"bufio"
	"fmt"
	"os"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/compiler"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/utils"
)

var running bool = false

func Start() bool {
	if running {
		logger.AddError("Shell is already running")
		return false
	}
	running = true
	AddShellCommands()
	return MainLoop()
}

func MainLoop() bool {
	for running {

		fmt.Print("gitgitgo>")

		in := bufio.NewReader(os.Stdin)

		line, err := in.ReadString('\n')

		if err != nil {
			logger.AddError("Error while reading input: " + err.Error())
			return false
		}

		compiler.ExecuteLine(line)

		logger.Log("")
	}

	return true
}

func Stop(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	running = false
	return true
}

func AddShellCommands() {

	compiler.AddSubcommand(cmdtool.NewSubcommand([]string{"exit"}, "Exit the shell", Stop))
	compiler.AddSubcommand(cmdtool.NewSubcommand([]string{"pwd"}, "Print the working directory", func(attValue []interface{}, fs *cmdtool.FlagSet) bool {
		dir, err := os.Getwd()
		if err != nil {
			logger.AddErrObj("Error while getting working directory", err)
			return false
		}
		fmt.Println(dir)
		return true
	}))
	compiler.AddSubcommand(cmdtool.NewSubcommand([]string{"clear"}, "Clear the terminal", func(attValue []interface{}, fs *cmdtool.FlagSet) bool {
		fmt.Print("\033[H\033[2J")
		return true
	}))
	cdCommand := cmdtool.NewSubcommand([]string{"cd"}, "Change the working directory", func(attValue []interface{}, fs *cmdtool.FlagSet) bool {
		if attValue == nil {
			logger.AddError("No path specified")
			return false
		}
		err := os.Chdir(attValue[0].(string))
		if err != nil {
			logger.AddErrObj("Error while changing directory", err)
			return false
		}
		utils.SetRepoName()
		return true
	})
	cdCommand.AddAttribute("path", "The path to change to")
	compiler.AddSubcommand(cdCommand)

	lsCommand := cmdtool.NewSubcommand([]string{"ls"}, "List the content in the directory", func(attValue []interface{}, fs *cmdtool.FlagSet) bool {

		ogPath, err := os.Getwd()
		if err != nil {
			logger.AddErrObj("Error while getting working directory", err)
			return false
		}
		if attValue[0] != nil {
			err := os.Chdir(attValue[0].(string))
			if err != nil {
				logger.AddErrObj("Error while changing directory", err)
				return false
			}
		}
		dir, err := os.Getwd()
		if err != nil {
			logger.AddErrObj("Error while getting working directory", err)
			ChangeToDir(ogPath)
			return false
		}
		files, err := os.ReadDir(dir)
		if err != nil {
			logger.AddErrObj("Error while reading directory", err)
			ChangeToDir(ogPath)
			return false
		}
		// print
		fmt.Println("")
		for _, file := range files {
			if file.IsDir() {
				fmt.Print("/")
			}
			fmt.Print(file.Name() + " ")
		}
		return ChangeToDir(ogPath)
	})
	lsCommand.AddAttribute("path", "The path to change to")
	compiler.AddSubcommand(lsCommand)

}

func ChangeToDir(dir string) bool {
	err := os.Chdir(dir)
	if err != nil {
		logger.AddErrObj("Error while changing directory", err)
		return false
	}
	return true
}
