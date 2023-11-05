package cmdtool

import (
	"fmt"
	"phil1436/GitGitGo-CLI/pkg/logger"
	"strings"
)

// FlagSet is a set of flags
type FlagSet struct {
	Flags []*Flag
}

// Creates a new FlagSet
func NewFlagSet() *FlagSet {
	return &FlagSet{}
}

// Adds a flag to the FlagSet
func (fs *FlagSet) AddFlag(flag *Flag) {
	fs.Flags = append(fs.Flags, flag)
}

// Get the value of a flag if it is defined otherwise nil
func (fs *FlagSet) GetValue(name string) interface{} {
	for _, flag := range fs.Flags {
		for _, n := range flag.Name {
			if n == name {
				return flag.Value
			}
		}
	}
	return nil
}

/* func (fs *FlagSet) GetStringValue(name string) interface{} {
	for _, flag := range fs.Flags {
		for _, n := range flag.Name {
			if n == name {
				return flag.Value.(string)
			}
		}
	}
	return nil
} */

// Check if a flag with the given name is defined
func (fs *FlagSet) IsDefined(name string) bool {
	for _, flag := range fs.Flags {
		for _, n := range flag.Name {
			if n == name {
				return true
			}
		}
	}
	return false
}

// Parse the given arguments and set the values of the flags
func (fs *FlagSet) Parse(args []string) {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg[0] == '-' {
			// This is a flag
			for _, flag := range fs.Flags {
				for _, name := range flag.Name {
					if strings.EqualFold(name, arg[1:]) {
						// This is the flag we are looking for
						if flag.BoolFlag {
							flag.Value = true
						} else {
							flag.Value = args[i+1]
							i++
						}
					}
				}
			}
		}
	}
}

// Check if all required flags are set
func (fs *FlagSet) IsFullFilled() bool {
	for _, flag := range fs.Flags {
		if flag.Value == nil {
			logger.AddError("Flag '-" + flag.Name[0] + "' is required")
			return false
		}
	}
	return true
}

// Concat two FlagSets. Returns a new FlagSet and does not modify the original ones!
func (fs *FlagSet) Concat(other *FlagSet) *FlagSet {
	result := NewFlagSet()
	for _, flag := range fs.Flags {
		result.AddFlag(flag)
	}
	for _, flag := range other.Flags {
		result.AddFlag(flag)
	}
	return result
}

// Convert the FlagSet to a string
func (fs *FlagSet) ToString() string {
	result := ""
	for _, flag := range fs.Flags {
		result += flag.ToString() + "\n"
	}
	return result
}

// Convert the FlagSet to a string with the current values
func (fs *FlagSet) GetStateString() string {
	if fs.IsEmpty() {
		return ""
	}
	result := "["
	for _, flag := range fs.Flags {
		if flag.Value == nil || flag.Value == "" || flag.Value == false {
			continue
		}
		result += flag.Name[0] + ": " + fmt.Sprintf("%v", flag.Value) + " | "
	}
	// remove last two characters
	result = result[:len(result)-3]
	return result + "]\n"
}

// Check if the FlagSet is empty
func (fs *FlagSet) IsEmpty() bool {
	return len(fs.Flags) == 0
}

func (fs *FlagSet) Copy() *FlagSet {
	newFlags := make([]*Flag, len(fs.Flags))
	copy(newFlags, fs.Flags)
	return &FlagSet{
		Flags: newFlags,
	}
}
