package cmdtool

// This represents a flag in the command line
type Flag struct {
	Name        []string
	Description string
	Value       interface{}
	BoolFlag    bool
}

// Creates a new flag
func NewFlag(name []string, description string, value interface{}, boolFlag bool) *Flag {
	return &Flag{
		Name:        name,
		Description: description,
		Value:       value,
		BoolFlag:    boolFlag,
	}
}

// makes a copy of the flag and returns it
func (f *Flag) Copy() *Flag {
	return NewFlag(f.Name, f.Description, f.Value, f.BoolFlag)
}

// ToString returns a string representation of the flag
func (f *Flag) ToString() string {
	result := "   "
	if len(f.Name) == 1 {
		result += "-" + f.Name[0]
	} else {
		result += "[-" + f.Name[0]
		i := 1
		for i < len(f.Name) {
			result += "|-" + f.Name[i]
			i++
		}
		result += "]"
	}

	result = result + " " + f.Description
	if f.Value != nil {
		result += " (optional)"
	}
	return result
}
