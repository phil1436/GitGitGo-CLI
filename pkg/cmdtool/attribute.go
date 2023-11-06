package cmdtool

type Attribute struct {
	Name        string
	Description string
}

func NewAttribute(name string, description string) *Attribute {
	return &Attribute{
		Name:        name,
		Description: description,
	}
}
