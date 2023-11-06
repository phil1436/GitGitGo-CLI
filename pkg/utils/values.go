package utils

import "strings"

const VERSION = "0.0.1"

var PROVIDER = ""
var GITHUBNAME = ""
var FULLNAME = ""
var REPONAME = ""

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
