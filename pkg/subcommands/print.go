package subcommands

import (
	"fmt"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/utils"
)

// Add a specified file to your project
func Print(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	logger.Log("*** gitgitgo PRINT ***")
	logger.Log(fs.GetStateString())
	if attValue[0] == nil {

		files := utils.GetAllFiles()
		if files == nil {
			return false
		}

		for _, file := range files {
			fmt.Println(file.ToString(fs.GetValue("name").(bool)))
		}
		return true
	}

	name := attValue[0].(string)

	if utils.IsRepoVarName(name) {
		fmt.Println("reponame: " + utils.REPONAME)
		return true
	}
	if utils.IsFullNameVarName(name) {
		fmt.Println("fullname: " + utils.FULLNAME)
		return true
	}
	if utils.IsGithubNameVarName(name) {
		fmt.Println("githubname: " + utils.GITHUBNAME)
		return true
	}
	if utils.IsProviderVarName(name) {
		fmt.Println("provider: " + utils.PROVIDER)
		return true
	}

	file := utils.GetFile(name)
	if file == nil {
		logger.AddError("File '" + name + "' not found")
		return false
	}

	fmt.Println(file.ToString(fs.GetValue("name").(bool)))

	return true
}
