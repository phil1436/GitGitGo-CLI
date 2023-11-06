package compiler

import (
	"fmt"
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

	if !activeSCmd.Run(args[start:], parsedGloablFlags) {
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

	if globalFlags.GetValue("dev").(bool) {
		// run in dev mode
		ExecuteLine("run setvalues.ggg")
	}

	if globalFlags.GetValue("quiet").(bool) {
		logger.Quiet()
	} else {
		logger.Unquiet()
	}

	if globalFlags.GetValue("provider").(string) != "" {
		newprovider := globalFlags.GetValue("provider").(string)
		if newprovider != utils.PROVIDER {
			utils.PROVIDER = newprovider
			logger.Log("Changed provider to '" + newprovider + "'")
			utils.ReloadFileManager()
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

// PrintVersion prints the current version
func PrintVersion(attValue []interface{}, fs *cmdtool.FlagSet) bool {
	fmt.Printf("GitGitGo-CLI version %s\n", utils.VERSION)
	return true
}

func AddSubcommand(scmd *cmdtool.Subcommand) {
	sbCmds = append(sbCmds, scmd)
}

func AddGlobalFlag(flag *cmdtool.Flag) {
	globalFlags.AddFlag(flag)
}
