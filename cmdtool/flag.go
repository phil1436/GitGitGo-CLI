package cmdtool

type Flag struct {
	Name        []string
	Description string
	Value       interface{}
	BoolFlag    bool
}

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
