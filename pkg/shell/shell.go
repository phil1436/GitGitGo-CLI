package shell

import (
	"bufio"
	"fmt"
	"os"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/compiler"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"strings"
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

		fmt.Print("gitgitgo> ")

		in := bufio.NewReader(os.Stdin)

		line, err := in.ReadString('\n')

		if err != nil {
			logger.AddError("Error while reading input: " + err.Error())
			return false
		}

		args := strings.Fields(line)

		compiler.ExecuteCommand(args, strings.EqualFold(args[0], "gitgitgo"))

		logger.Log("")
	}

	return true
}

func Stop(fs *cmdtool.FlagSet) bool {
	running = false
	return true
}

func AddShellCommands() {

	compiler.AddSubcommand(cmdtool.NewSubcommand([]string{"exit"}, "Exit the shell", Stop))

}
