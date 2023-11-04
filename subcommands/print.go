package subcommands

import (
	"fmt"
	"phil1436/GitGitGo-CLI/cmdtool"
	"phil1436/GitGitGo-CLI/utils"
)

// Add a specified file to your project
func Print(fs *cmdtool.FlagSet) bool {
	fmt.Println("gitgitgo print")
	fmt.Println(fs.GetStateString())

	if fs.GetValue("file") == "" {
		for _, file := range utils.GetAllFiles(utils.GITHUBNAME) {
			fmt.Println(file.ToString(fs.GetValue("name").(bool)))
		}
		return true
	}

	file := utils.GetFile(fs.GetValue("file").(string), utils.GITHUBNAME)
	if file == nil {
		fmt.Println("File '" + fs.GetValue("file").(string) + "' not found")
		return false
	}

	fmt.Println(file.ToString(fs.GetValue("name").(bool)))

	return true
}
