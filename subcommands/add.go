package subcommands

import (
	"fmt"
	"phil1436/GitGitGo-CLI/cmdtool"
	"phil1436/GitGitGo-CLI/utils"
)

// Add a specified file to your project
func Add(fs *cmdtool.FlagSet) bool {
	fmt.Println("gitgitgo add")
	fmt.Println(fs.GetStateString())

	file := utils.GetFile(fs.GetValue("file").(string), utils.GITHUBNAME)
	if file == nil {
		fmt.Println("File '" + fs.GetValue("file").(string) + "' not found")
		return false
	}

	return file.Save(fs.GetValue("destination").(string), utils.GITHUBNAME, utils.FULLNAME, utils.REPONAME)
}
