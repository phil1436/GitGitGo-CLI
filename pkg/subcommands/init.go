package subcommands

import (
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/utils"
	"strings"
)

func Init(attValue []interface{}, fs *cmdtool.FlagSet) bool {
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

	withKeywords := attValue[0] != nil
	keywords := []string{""}
	if withKeywords {
		keywords = strings.Split(attValue[0].(string), ",")
	}

	ignore := strings.Split(fs.GetValue("ignore").(string), ",")

	gitIgnore := utils.GetFile(".gitignore")

	for _, file := range files {
		// ignore .gitignore
		if file.Name == ".gitignore" {
			continue
		}
		if withKeywords {
			if !file.ContainsKeywordArr(keywords) {
				continue
			}

		} else {
			if !file.OnInit {
				continue
			}
		}
		if utils.ArrContains(ignore, file.Name) {
			logger.Log("Ignore file: " + file.Name)
			continue
		}

		// add file to gitignore
		if file.Ignore && gitIgnore != nil {
			gitIgnore.Text += "\n" + file.Name
		}

		if !file.Save(fs.GetValue("destination").(string), utils.GITHUBNAME, utils.FULLNAME, utils.REPONAME, "", fs.GetValue("force").(bool), fs.GetValue("dryrun").(bool)) {
			result = false
			continue
		}
	}

	if gitIgnore != nil {
		gitIgnore.Save(fs.GetValue("destination").(string), utils.GITHUBNAME, utils.FULLNAME, utils.REPONAME, ".gitignore", fs.GetValue("force").(bool), fs.GetValue("dryrun").(bool))
	}

	return result
}
