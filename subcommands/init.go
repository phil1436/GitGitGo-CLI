package subcommands

import (
	"GitGitGo-CLI/cmdtool"
	"fmt"
)

func Init(fs *cmdtool.FlagSet) {
	fmt.Println("init")
	fmt.Println(fs.GetStateString())
}
