package utils

import (
	"phil1436/GitGitGo-CLI/pkg/logger"
	"strings"
)

var PROVIDER = ""
var GITHUBNAME = ""
var FULLNAME = ""
var REPONAME = ""

var BINARY = ""
var HOME = ""

// Set to the provider and reload the file manager with the new provider
func SetProvider(newProvider string) bool {
	if newProvider == "" {
		logger.AddError("Empty provider")
		return false
	}
	if newProvider == PROVIDER {
		return true
	}

	oldProvider := PROVIDER
	PROVIDER = newProvider
	if !ReloadFileManager() {
		if newProvider != oldProvider {
			logger.AddError("Provider '" + newProvider + "' not found! Changed back to '" + oldProvider + "'")
		} else {
			logger.AddError("Provider '" + newProvider + "' not found!")
		}
		PROVIDER = oldProvider
		return false
	}
	return true
}

func IsProviderVarName(name string) bool {
	return strings.EqualFold(name, "provider") || strings.EqualFold(name, "p")
}

func IsGithubNameVarName(name string) bool {
	return strings.EqualFold(name, "githubname") || strings.EqualFold(name, "gname") || strings.EqualFold(name, "gn")
}

func IsFullNameVarName(name string) bool {
	return strings.EqualFold(name, "fullname") || strings.EqualFold(name, "fname") || strings.EqualFold(name, "fn")
}
func IsRepoVarName(name string) bool {
	return strings.EqualFold(name, "reponame") || strings.EqualFold(name, "rname") || strings.EqualFold(name, "rn")
}
