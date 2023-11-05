package subcommands

import (
	"fmt"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/utils"
)

// Add a specified file to your project
func Print(fs *cmdtool.FlagSet) bool {
	logger.Log("*** gitgitgo PRINT ***")
	logger.Log(fs.GetStateString())

	if fs.GetValue("file") == "" {

		files := utils.GetAllFiles()
		if files == nil {
			return false
		}

		for _, file := range files {
			fmt.Println(file.ToString(fs.GetValue("name").(bool)))
		}
		return true
	}

	file := utils.GetFile(fs.GetValue("file").(string))
	if file == nil {
		fmt.Println("File '" + fs.GetValue("file").(string) + "' not found")
		return false
	}

	fmt.Println(file.ToString(fs.GetValue("name").(bool)))

	return true
}
