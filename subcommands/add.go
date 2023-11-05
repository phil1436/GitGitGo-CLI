package subcommands

import (
	"path/filepath"
	"phil1436/GitGitGo-CLI/cmdtool"
	"phil1436/GitGitGo-CLI/logger"
	"phil1436/GitGitGo-CLI/utils"
)

// Add a specified file to your project
func Add(fs *cmdtool.FlagSet) bool {
	logger.Log("*** GitGitGo ADD ***\n")
	logger.Log(fs.GetStateString())

	if fs.GetValue("dryrun").(bool) {
		logger.Log("DRYRUN: No files will be created\n")
	}

	file := utils.GetFile(fs.GetValue("file").(string))
	if file == nil {
		logger.AddError("File '" + fs.GetValue("file").(string) + "' not found")
		return false
	}

	reponame := utils.REPONAME
	destination := fs.GetValue("destination").(string)

	if destination != "." {
		reponame = filepath.Base(destination)
	}

	return file.Save(destination, utils.GITHUBNAME, utils.FULLNAME, reponame, fs.GetValue("as").(string), fs.GetValue("force").(bool), fs.GetValue("dryrun").(bool))
}
