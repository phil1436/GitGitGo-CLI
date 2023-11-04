package subcommands

import (
	"fmt"
	"phil1436/GitGitGo-CLI/cmdtool"
	"phil1436/GitGitGo-CLI/utils"
)

func Init(fs *cmdtool.FlagSet) bool {
	fmt.Println("gitgitgo init")
	fmt.Println(fs.GetStateString())

	for _, file := range utils.GetAllFiles(utils.GITHUBNAME) {
		if !file.OnInit {
			continue
		}
		if !file.Save(fs.GetValue("destination").(string), utils.GITHUBNAME, utils.FULLNAME, utils.REPONAME) {
			fmt.Println("File '" + file.Name + "' could not be created")
			return false
		}
		fmt.Println("File '" + file.Name + "' created...")

	}

	return true
}
