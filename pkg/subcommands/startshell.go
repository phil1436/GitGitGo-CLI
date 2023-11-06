package subcommands

import (
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/shell"
)

func StartShell(fs *cmdtool.FlagSet) bool {
	logger.Log("*** gitgitgo SHELL ***")
	logger.Log(fs.GetStateString())

	return shell.Start()
}