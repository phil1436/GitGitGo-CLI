package subcommands

import (
	"path/filepath"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/utils"
)

// Add a specified file to your project
func Add(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	logger.Log("*** GitGitGo ADD ***\n")
	logger.Log(fs.GetStateString())

	if attValue == nil {
		logger.AddError("No file specified")
		return false
	}

	logger.Log("Add file: " + attValue[0].(string))

	if fs.GetValue("dryrun").(bool) {
		logger.Log("DRYRUN: No files will be created\n")
	}
	fileName := attValue[0].(string)
	file := utils.GetFile(fileName)
	if file == nil {
		logger.AddError("File '" + fileName + "' not found")
		return false
	}

	reponame := utils.REPONAME
	destination := fs.GetValue("destination").(string)

	if destination != "." {
		reponame = filepath.Base(destination)
	}

	// Add to gitignore
	file.AddToGitIgnore()

	return file.Save(destination, utils.GITHUBNAME, utils.FULLNAME, reponame, fs.GetValue("as").(string), fs.GetValue("force").(bool), fs.GetValue("dryrun").(bool))
}
