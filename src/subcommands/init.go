package subcommands

import (
	"phil1436/GitGitGo-CLI/src/cmdtool"
	"phil1436/GitGitGo-CLI/src/logger"
	"phil1436/GitGitGo-CLI/src/utils"
)

func Init(fs *cmdtool.FlagSet) bool {
	logger.Log("*** GitGitGo INIT ***\n")
	logger.Log(fs.GetStateString())

	if fs.GetValue("dryrun").(bool) {
		logger.Log("DRYRUN: No files will be created\n")
	}

	result := true

	files := utils.GetAllFiles()
	if files == nil {
		return false
	}

	for _, file := range files {
		if !file.OnInit {
			continue
		}
		if !file.Save(fs.GetValue("destination").(string), utils.GITHUBNAME, utils.FULLNAME, utils.REPONAME, "", fs.GetValue("force").(bool), fs.GetValue("dryrun").(bool)) {
			result = false
			continue
		}
	}

	return result
}
