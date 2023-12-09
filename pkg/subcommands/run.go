package subcommands

import (
	"os"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/compiler"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"strings"
)

// Run a specified file
func Run(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	logger.Log("*** GitGitGo RUN ***\n")
	logger.Log(fs.GetStateString())

	if attValue[0] == nil {
		logger.AddError("No file given")
		return false
	}

	logger.Log("Run file: " + attValue[0].(string))

	// read file
	fileContent, err := os.ReadFile(attValue[0].(string))
	if err != nil {
		logger.AddErrObj("Failed to read file", err)
		return false
	}

	// split file into lines
	lines := strings.Split(string(fileContent), "\n")

	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "#") || strings.TrimSpace(line) == "" {
			continue
		}
		if !compiler.ExecuteLine(line) {
			logger.AddError("Failed to execute command '" + line + "'")
			return false
		}
	}

	return true
}
