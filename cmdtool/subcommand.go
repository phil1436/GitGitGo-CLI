package cmdtool

// class Subcommand
//
// This class is used to create a subcommand for the CLI tool.
// It contains the following fields:
// - Name: the name of the subcommand
// - Description: a description of the subcommand
// - FlagSet: a FlagSet object that contains the flags for the subcommand
// - Handler: a function that is called when the subcommand is executed
type Subcommand struct {
	Names       []string
	Description string
	Handler     func(fs *FlagSet) bool
	FlagSet     *FlagSet
}

func NewSubcommand(names []string, description string, handler func(fs *FlagSet) bool) *Subcommand {
	return &Subcommand{
		Names:       names,
		Description: description,
		Handler:     handler,
		FlagSet:     NewFlagSet(),
	}
}

func (s *Subcommand) AddFlag(flag *Flag) {
	s.FlagSet.AddFlag(flag)
}

func (s *Subcommand) Run(args []string, fs *FlagSet) bool {
	s.FlagSet.Parse(args)
	if !s.FlagSet.IsFullFilled() {
		return false
	}
	x := s.FlagSet
	if fs != nil {
		fs.Parse(args)
		if !fs.IsFullFilled() {
			return false
		}
		x = s.FlagSet.Concat(fs)
	}
	return s.Handler(x)
}

func (s *Subcommand) ToString() string {
	result := ""
	if len(s.Names) == 1 {
		result += s.Names[0]
	} else {
		result += "[" + s.Names[0]
		//iterate over the rest of the names
		i := 1
		for i < len(s.Names) {
			result += "|" + s.Names[i]
			i++
		}
		result += "]"
	}

	if s.FlagSet.IsEmpty() {
		return result + ": " + s.Description + "\n"
	}
	result = result + " [flags]: " + s.Description
	result += "\n" + s.FlagSet.ToString()
	return result
}
