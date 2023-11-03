package cmdtool

import "strings"

type FlagSet struct {
	Flags []*Flag
}

func NewFlagSet() *FlagSet {
	return &FlagSet{}
}

func (fs *FlagSet) AddFlag(flag *Flag) {
	fs.Flags = append(fs.Flags, flag)
}

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

func (fs *FlagSet) Parse(args []string) {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg[0] == '-' {
			// This is a flag
			for _, flag := range fs.Flags {
				for _, name := range flag.Name {
					if strings.EqualFold(name, arg[1:]) {
						// This is the flag we are looking for
						flag.Value = args[i+1]
						i++
					}
				}
			}
		}
	}
}

func (fs *FlagSet) IsFullFilled(msg *string) bool {
	for _, flag := range fs.Flags {
		if flag.Value == nil {

			*msg = "Flag '-" + flag.Name[0] + "' is required"
			return false
		}
	}
	return true
}

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

func (fs *FlagSet) ToString() string {
	result := ""
	for _, flag := range fs.Flags {
		result += flag.ToString() + "\n"
	}
	return result
}
func (fs *FlagSet) GetStateString() string {
	result := ""
	for _, flag := range fs.Flags {
		result += flag.Name[0] + ": " + flag.Value.(string) + "\n"
	}
	return result
}

func (fs *FlagSet) IsEmpty() bool {
	return len(fs.Flags) == 0
}
