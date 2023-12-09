package compiler

import (
	"fmt"
	"os"
	"phil1436/GitGitGo-CLI/pkg/cmdtool"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"phil1436/GitGitGo-CLI/pkg/utils"
	"strings"
)

var sbCmds []*cmdtool.Subcommand
var globalFlags *cmdtool.FlagSet = cmdtool.NewFlagSet()

// ExecuteCommand executes the command
func ExecuteCommand(args []string, ignoreFirstArg bool) bool {
	activeSCmd := GetActiveSubcommand(args, ignoreFirstArg)

	if activeSCmd == nil {
		//GlobalHelp(nil, nil)
		return false
	}
	start := 1
	if ignoreFirstArg {
		start = 2
	}

	parsedGloablFlags := globalFlags.Copy()
	parsedGloablFlags.Parse(args[start:])
	BeforeRun(parsedGloablFlags)
	runSuccess := activeSCmd.Run(args[start:], parsedGloablFlags)
	if !runSuccess || logger.IsError() {
		fmt.Println("\nAn Error occured while executing the command '" + activeSCmd.Names[0] + "':")
		logger.PrintErrors()
		fmt.Println("\nUse 'gitgitgo help' to get help")
		return false
	}
	return true
}

func ExecuteLine(line string) bool {
	args := strings.Fields(line)
	for i, arg := range args {
		if strings.HasPrefix(arg, "\"") {
			args[i] = strings.TrimPrefix(arg, "\"")
			for j := i + 1; j < len(args); j++ {
				if strings.HasSuffix(args[j], "\"") {
					args[i] += " " + strings.TrimSuffix(args[j], "\"")
					args = append(args[:j], args[j+1:]...)
					break
				} else {
					args[i] += " " + args[j]
					args = append(args[:j], args[j+1:]...)
					j--
				}
			}
		}
	}
	return ExecuteCommand(args, strings.EqualFold(args[0], "gitgitgo"))
}

// BeforeRun is called before the subcommand is executed
func BeforeRun(globalFlags *cmdtool.FlagSet) {
	HandleGitGitGoConfigFile(".")

	if globalFlags.GetValue("dev").(bool) {
		// run in dev mode
		err := os.Chdir("dev")
		if err != nil {
			logger.AddErrObj("Error while changing directory", err)
		}
	}

	if globalFlags.GetValue("quiet").(bool) {
		logger.Quiet()
	} else {
		logger.Unquiet()
	}

	if globalFlags.GetValue("provider").(string) != "" {
		newprovider := globalFlags.GetValue("provider").(string)
		if utils.SetProvider(newprovider) {
			logger.Log("Changed provider to '" + newprovider + "'")
		}
	}
	if globalFlags.GetValue("githubname").(string) != "" {
		utils.GITHUBNAME = globalFlags.GetValue("githubname").(string)
	}
	if globalFlags.GetValue("fullname").(string) != "" {
		utils.FULLNAME = globalFlags.GetValue("fullname").(string)
	}
	if globalFlags.GetValue("reponame").(string) != "" {
		utils.REPONAME = globalFlags.GetValue("reponame").(string)
	}

}

func HandleGitGitGoConfigFile(directory string) {
	// check if file exists

	text, err := os.ReadFile(directory + "/.gitgitgo")
	if err != nil {
		text, err = os.ReadFile(directory + "/.gitgitgoc")
		if err != nil {
			return
		}
	}

	lines := strings.Split(string(text), "\n")
	seperator := "="
	for i, line := range lines {

		// get seperator
		if i == 0 && strings.HasPrefix(strings.TrimSpace(line), "#") {
			if strings.TrimSpace(line) == "#:" || strings.TrimSpace(line) == "# :" {
				seperator = ":"
				continue
			}
			if strings.TrimSpace(line) == "#-" || strings.TrimSpace(line) == "# -" {
				seperator = "-"
				continue
			}
			if strings.TrimSpace(line) == "#=" || strings.TrimSpace(line) == "# =" {
				seperator = "="
				continue
			}
		}

		// ignore empty lines and comments
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}

		keyval := strings.Split(line, seperator)
		if len(keyval) != 2 {
			logger.AddError("Invalid line in .gitgitgoc file: '" + line + "'")
			return
		}

		key := strings.TrimSpace(keyval[0])
		val := strings.TrimSpace(keyval[1])

		if utils.IsProviderVarName(key) {
			utils.SetProvider(val)
		}
		if utils.IsGithubNameVarName(key) {
			utils.GITHUBNAME = val
		}
		if utils.IsFullNameVarName(key) {
			utils.FULLNAME = val
		}
		if utils.IsRepoVarName(key) {
			utils.REPONAME = val
		}
	}

}

// GetActiveSubcommand returns the active subcommand
func GetActiveSubcommand(args []string, ignoreFirstArg bool) *cmdtool.Subcommand {
	// Get the subcommand name
	first := 0

	if ignoreFirstArg {
		first = 1
	}

	if len(args) < (first + 1) {
		fmt.Println("No command given")
		fmt.Println("")
		fmt.Println("Use 'gitgitgo help' to get help")
		return nil
	}

	for _, sb := range sbCmds {
		for _, name := range sb.Names {
			if strings.EqualFold(name, args[first]) {
				return sb
			}
		}
	}
	fmt.Println("Unknown command '" + args[first] + "'\n")
	fmt.Println("Use 'gitgitgo help' to get help")

	return nil
}

func Help(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	if attValue[0] == nil {
		GlobalHelp(nil, nil)
		return true
	}
	for _, sb := range sbCmds {
		for _, name := range sb.Names {
			if name == attValue[0].(string) {
				fmt.Println(name + " HELP:\n")
				fmt.Println(sb.ToString())
				return true
			}
		}
	}
	logger.AddError("Unknown command '" + attValue[0].(string) + "'\n")
	return false
}

// GlobalHelp prints the global help
// Set fs to nil!
func GlobalHelp(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	fmt.Println("GitGitGo-CLI - a CLI tool to help you manage your git repositories. ")
	fmt.Println("")
	//fmt.Printf("Usage: %s [command] [flags]\n", os.Args[0])
	fmt.Println("Usage: gitgitgo [command] [flags]")
	fmt.Println("")
	fmt.Println("Commands:")
	for _, sb := range sbCmds {
		fmt.Println(sb.ToString())
	}
	fmt.Println("")
	if !globalFlags.IsEmpty() {
		fmt.Println("Global Flags (Use with every command):")
		fmt.Println(globalFlags.ToString())
	}
	return true
}

func AddSubcommand(scmd *cmdtool.Subcommand) {
	sbCmds = append(sbCmds, scmd)
}

func AddGlobalFlag(flag *cmdtool.Flag) {
	globalFlags.AddFlag(flag)
}
