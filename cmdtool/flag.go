package cmdtool

type Flag struct {
	Name        []string
	Description string
	Value       interface{}
}

func (f *Flag) ToString() string {
	result := "   "
	if len(f.Name) == 1 {
		result += "-" + f.Name[0]
	} else {
		result += "[-" + f.Name[0]
		//iterate over the rest of the names
		i := 1
		for i < len(f.Name) {
			result += "|-" + f.Name[i]
			i++
		}
		result += "]"
	}

	return result + " " + f.Description
}
