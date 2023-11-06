package cmdtool

import (
	"strings"
)

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
	Handler     func(attValue []interface{}, fs *FlagSet) bool
	FlagSet     *FlagSet
	Attribute   []*Attribute
}

// Creates a new subcommand
func NewSubcommand(names []string, description string, handler func(attValue []interface{}, fs *FlagSet) bool) *Subcommand {
	return &Subcommand{
		Names:       names,
		Description: description,
		Handler:     handler,
		FlagSet:     NewFlagSet(),
	}
}

// Adds a flag to the subcommand
func (s *Subcommand) AddFlag(flag *Flag) {
	s.FlagSet.AddFlag(flag)
}

// Adds an attribute to the subcommand
func (s *Subcommand) AddAttribute(name string, description string) {
	s.Attribute = append(s.Attribute, NewAttribute(name, description))
}

// Run runs the subcommand
func (s *Subcommand) Run(args []string, fs *FlagSet) bool {
	myFlagSet := s.FlagSet.Copy()
	myFlagSet.Parse(args)
	if !myFlagSet.IsFullFilled() {
		return false
	}
	x := myFlagSet
	if fs != nil {
		fs.Parse(args)
		if !fs.IsFullFilled() {
			return false
		}
		x = x.Concat(fs)
	}
	if s.Attribute != nil {
		return s.Handler(s.ParseAttributes(args), x)
	}
	return s.Handler(nil, x)
}

func (s *Subcommand) ParseAttributes(args []string) []interface{} {
	if s.Attribute == nil {
		return make([]interface{}, 0)
	}
	result := make([]interface{}, len(s.Attribute))
	flagDetected := false
	for i := range s.Attribute {
		if len(args) <= i || flagDetected {
			result[i] = nil
			continue
		}
		if strings.HasPrefix(args[i], "-") {
			result[i] = nil
			flagDetected = true
			continue
		}
		if strings.HasPrefix(args[i], "\\-") {
			result[i] = args[i][1:]
			continue
		}
		result[i] = args[i]
	}
	return result
}

// ToString returns a string representation of the subcommand
func (s *Subcommand) ToString() string {
	result := ""
	if len(s.Names) == 1 {
		result += s.Names[0]
	} else {
		result += "[" + s.Names[0]
		i := 1
		for i < len(s.Names) {
			result += "|" + s.Names[i]
			i++
		}
		result += "]"
	}

	if s.Attribute != nil {
		for _, att := range s.Attribute {
			result += " <" + att.Name + ">"
		}
	}

	if s.FlagSet.IsEmpty() {
		result = result + ": " + s.Description
	} else {
		result = result + " <flags>: " + s.Description
	}
	if s.Attribute != nil {
		for _, att := range s.Attribute {
			result += "\n   <" + att.Name + ">: " + att.Description
		}
	}
	result += "\n" + s.FlagSet.ToString()
	return result
}
