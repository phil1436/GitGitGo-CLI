package subcommands

import (
	"GitGitGo-CLI/cmdtool"
	"fmt"
)

func Add(fs *cmdtool.FlagSet) {
	fmt.Println("add")
	fmt.Println(fs.GetStateString())
}
